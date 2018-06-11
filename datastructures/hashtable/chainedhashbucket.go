package hashtable

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	// "log"
	"sync"
	"time"
)

const (
	// MaxCapacity is the maximum capacity for the hash table. beyond that, there will be no more resizing. chosen arbitrarily.
	MaxCapacity uint64 = 4294967311
)

// ChainedHashBucket is a thread-safe hash table implementation.
// For collusions (elements that are mapped by the hash function to the same slot of the table), an ordered doubly linked list is used.
type ChainedHashBucket struct {
	rwlock                 *sync.RWMutex
	data                   nodes  // this is where the data resides
	resizeData             nodes  // if an incremental resize is in progress, then this is temporary table for it.
	numItems               uint64 // number of key value pairs in the table. incremented with a put and decremented with a logical delete
	numPhysicalNodes       uint64 // number of pysical key value pairs in the table. incremented with a put but not affected by a logical delete
	resizeNumItems         uint64
	resizeNumPhysicalNodes uint64
	key0                   uint64 // first 64 bit key for siphash
	key1                   uint64 // second 64 bit key for siphash
	isResizing             bool   // true if the  bucket being resized
	resizeInd              uint64 // resize index
}

type nodes []*kvnode

// NewChainedHashBucket returns a new ChainedHashBucket with initial capacity of 'capacity'
func NewChainedHashBucket(capacity uint64) *ChainedHashBucket {
	key0, key1, err := generateSipKeys()
	if err != nil {
		return nil
	}
	if capacity <= 17 {
		capacity = 17
	} else if capacity > MaxCapacity {
		capacity = MaxCapacity
	} else {
		for !isPrime(capacity) {
			capacity++
		}
		if capacity > MaxCapacity {
			capacity = MaxCapacity
		}
	}
	// log.Println("NewChainedHashBucket - capacity:", capacity)
	b := &ChainedHashBucket{
		rwlock:     new(sync.RWMutex),
		data:       nodes(make([]*kvnode, capacity)),
		resizeData: nil,
		key0:       key0,
		key1:       key1,
	}
	return b
}

// LoadFactor is n / m where n is the number of physical nodes (key value pairs, logically deleted or not)
// in the table and m is the capacity of the table.
func (t *ChainedHashBucket) LoadFactor() float64 {
	t.rwlock.RLock()
	r := t.loadFactor()
	t.rwlock.RUnlock()
	return r
}

func (t *ChainedHashBucket) loadFactor() float64 {
	return float64(t.numPhysicalNodes) / float64(len(t.data))
}

// Length is the number of existing elements in the bucket
func (t *ChainedHashBucket) Length() uint64 {
	t.rwlock.RLock()
	r := t.resizeNumItems + t.numItems
	t.rwlock.RUnlock()
	return r
}

// Load is the number of physical nodes in the bucket. It may include nodes that have been logically deleted
func (t *ChainedHashBucket) Load() uint64 {
	t.rwlock.RLock()
	r := t.resizeNumPhysicalNodes + t.numPhysicalNodes
	t.rwlock.RUnlock()
	return r
}

// Capacity returns the length of underlying data slice in the bucket.
func (t *ChainedHashBucket) Capacity() uint64 {
	t.rwlock.RLock()
	r := uint64(len(t.data))
	t.rwlock.RUnlock()
	return r
}

// ResizePartial resizes the hash bucket for a limited duration.
// When resize duration exceeds the input duration, resize will stop as soon as possible.
// Returns an error if there is a reason not to start the resize.
// Also returns true if resize is complete,
func (t *ChainedHashBucket) ResizePartial(dur time.Duration, targetLoadFactor float64) (bool, error) {
	t.rwlock.Lock()
	defer t.rwlock.Unlock()

	if t.isResizing {
		return false, fmt.Errorf("a resize is already in progress")
	}
	if uint64(len(t.data)) >= MaxCapacity {
		return false, fmt.Errorf("failed to resize a bucket at max capacity")
	}
	t.isResizing = true
	timer := time.NewTimer(dur)

	// log.Println("partial resizing lock acquired, target duration", dur)
	// t0 := time.Now()
	if t.resizeData == nil {
		if t.loadFactor() <= targetLoadFactor {
			t.isResizing = false
			return true, nil
		}
		capacity := uint64(2)
		for float64(t.numItems)/float64(capacity) > targetLoadFactor {
			capacity *= 2
			if capacity > MaxCapacity {
				capacity = MaxCapacity
				break
			}
		}

		for !isPrime(capacity) {
			capacity++
		}
		if capacity > MaxCapacity {
			capacity = MaxCapacity
		}
		t.resizeInd = 0
		t.resizeData = nodes(make([]*kvnode, capacity))
	} else {
		// log.Println("previous incomplete partial resize detected, target load factor will be ignored")
	}

	// t is already locked, there might not be a need to lock resizeBucket as well.
	// however, let locking remain for the moment.
	// t.resizeBucket.rwlock.Lock()
	// defer t.resizeBucket.rwlock.Unlock()
	var convertedItemCount uint64

	for {
		select {
		case <-timer.C:
			t.isResizing = false
			timer.Stop()
			// log.Println("partial resizing expired, resizeInd", t.resizeInd, "converteditemcount", convertedItemCount)
			if t.resizeInd >= uint64(len(t.data)) {
				t.data = t.resizeData
				t.resizeData = nil
				t.resizeInd = 0
				t.isResizing = false
				t.numItems = t.resizeNumItems
				t.numPhysicalNodes = t.resizeNumPhysicalNodes
				t.resizeNumItems = 0
				t.resizeNumPhysicalNodes = 0

				// log.Println("partial resizing completed", time.Since(t0), "convertedItemCount", convertedItemCount, "new capacity", len(t.data))
			}
			return t.resizeInd >= uint64(len(t.data)), nil
		default:
			if t.resizeInd >= uint64(len(t.data)) {
				t.data = t.resizeData
				t.resizeData = nil
				t.resizeInd = 0
				t.isResizing = false
				t.numItems = t.resizeNumItems
				t.numPhysicalNodes = t.resizeNumPhysicalNodes
				t.resizeNumItems = 0
				t.resizeNumPhysicalNodes = 0

				// log.Println("partial resizing completed", time.Since(t0), "convertedItemCount", convertedItemCount, "new capacity", len(t.data))

				return true, nil
			}

			// log.Println("partial resizing resizeInd", t.resizeInd)
			n := t.data[t.resizeInd]
			for n != nil {
				// logically deleted item, skip
				if n.meta&(1<<7) != 0 {
					n = n.next
					continue
				}
				loc := Hash(t.key0, t.key1, n.key) % uint64(len(t.resizeData))

				switch rc := t.resizeData.put(n.key, n.value, loc, false, false); rc {
				case 0:
					t.resizeNumItems++
					t.resizeNumPhysicalNodes++
					t.numItems--
				case 2:
					t.resizeNumItems++
				}

				n = n.next
				convertedItemCount++
			}
			t.resizeInd++
		}
	}

}

// ResizeFull fully resizes the hash bucket with a target load factor taken into consideration.
func (t *ChainedHashBucket) ResizeFull(targetLoadFactor float64) error {
	t.rwlock.Lock()
	defer t.rwlock.Unlock()

	if t.isResizing {
		return fmt.Errorf("a resize is already in progress")
	}
	if uint64(len(t.data)) >= MaxCapacity {
		return fmt.Errorf("failed to resize a bucket at max capacity")
	}
	t.isResizing = true

	// log.Println("full resizing lock acquired")
	// t0 := time.Now()
	if t.resizeData == nil {

		if t.loadFactor() <= targetLoadFactor {
			t.isResizing = false
			return nil
		}
		// use number of items instead of physical nodes because logically deleted nodes will not be copied anyway.
		capacity := uint64(2)
		for float64(t.numItems)/float64(capacity) > targetLoadFactor {
			capacity *= 2
			if capacity > MaxCapacity {
				capacity = MaxCapacity
				break
			}
		}

		for !isPrime(capacity) {
			capacity++
		}
		if capacity > MaxCapacity {
			capacity = MaxCapacity
		}
		t.resizeInd = 0
		t.resizeData = nodes(make([]*kvnode, capacity))
	}

	var convertedItemCount uint64

	for {
		if t.resizeInd >= uint64(len(t.data)) {
			t.data = t.resizeData
			t.resizeData = nil
			t.resizeInd = 0
			t.isResizing = false
			t.numItems = t.resizeNumItems
			t.numPhysicalNodes = t.resizeNumPhysicalNodes
			t.resizeNumItems = 0
			t.resizeNumPhysicalNodes = 0

			// log.Println("full resizing completed", time.Since(t0), "convertedItemCount", convertedItemCount, "new capacity", len(t.data))

			return nil
		}

		// log.Println("partial resizing resizeInd", t.resizeInd)
		n := t.data[t.resizeInd]
		for n != nil {
			// logically deleted item, skip
			if n.meta&(1<<7) != 0 {
				n = n.next
				continue
			}
			loc := Hash(t.key0, t.key1, n.key) % uint64(len(t.resizeData))

			switch rc := t.resizeData.put(n.key, n.value, loc, false, false); rc {
			case 0:
				t.resizeNumItems++
				t.resizeNumPhysicalNodes++
				t.numItems--
			case 2:
				t.resizeNumItems++
			}

			n = n.next
			convertedItemCount++
		}
		t.resizeInd++
	}

}

// Delete soft deletes a key if it exists, returns an error if it does not.
func (t *ChainedHashBucket) Delete(k []byte) error {
	t.rwlock.Lock()
	defer t.rwlock.Unlock()

	hval := Hash(t.key0, t.key1, k) % uint64(len(t.data))
	var node *kvnode

	if t.resizeData != nil && hval < t.resizeInd {
		hval = Hash(t.key0, t.key1, k) % uint64(len(t.resizeData))
		node = t.resizeData[hval]
	} else {
		node = t.data[hval]
	}

	for node != nil {
		c := bytes.Compare(k, node.key)
		switch c {
		case -1:
			return fmt.Errorf("key %v not found in bucket", string(k))
		case 0:
			if node.meta&(1<<7) != 0 {
				return fmt.Errorf("node is logically deleted key %v not found in bucket", string(k))
			}
			node.meta |= 1 << 7
			t.numItems--
			return nil
		case 1:
			node = node.next
		}
	}

	return fmt.Errorf("key %v not found in bucket", string(k))
}

// Put serves as insert or update a key k, with the given value v
func (t *ChainedHashBucket) Put(k []byte, v []byte) error {
	t.rwlock.Lock()
	defer t.rwlock.Unlock()

	var loc uint64

	var target nodes

	if t.resizeData != nil {
		// log.Println("putting", string(k), "to resize bucket")
		loc = Hash(t.key0, t.key1, k) % uint64(len(t.resizeData))
		target = t.resizeData
	} else {
		// log.Println("putting", string(k), "to bucket")
		loc = Hash(t.key0, t.key1, k) % uint64(len(t.data))
		target = t.data
	}

	switch rc := target.put(k, v, loc, true, true); rc {
	case 0:
		t.numItems++
		t.numPhysicalNodes++
	case 2:
		t.numItems++
	}
	return nil
}

// Get retrieves value associated with given key. Returns nil if the key does not exist.
func (t *ChainedHashBucket) Get(k []byte) []byte {
	t.rwlock.RLock()
	defer t.rwlock.RUnlock()
	hval := Hash(t.key0, t.key1, k) % uint64(len(t.data))
	var node *kvnode

	if t.resizeData != nil && hval < t.resizeInd {
		// log.Println("Getting key", string(k), "from resize bucket")
		hval = Hash(t.key0, t.key1, k) % uint64(len(t.resizeData))
		node = t.resizeData[hval]
	} else {
		// log.Println("Getting key", string(k), "from standard bucket")
		node = t.data[hval]
	}

	for node != nil {
		c := bytes.Compare(k, node.key)
		switch c {
		case 1:
			node = node.next
			continue
		case 0:
			if node.meta&(1<<7) != 0 {
				return nil
			}
			return node.value
		case -1:
			return nil
		}
	}

	return nil
}

// put puts the key value pair to the input named type slice.
// with the following return codes we can keep track of number of items and number of physical nodes in a bucket
// rc (return code):
//      -1 if no data is written
//       0 if a new node is created
//       1 if an override occurs on an existing key
//       2 if an overide occurs on a logically deleted node
func (n nodes) put(k, v []byte, loc uint64, oe, od bool) (rc int) {
	node := n[loc]

	for node != nil {
		c := bytes.Compare(k, node.key)
		switch c {
		case 1:
			if node.next != nil {
				node = node.next
				continue
			}
			kvn := newKVNode(k, v, 0, node, nil)
			node.next = kvn
			rc = 0
			return
		case 0:
			rc = -1
			if node.meta&(1<<7) != 0 {
				if od {
					rc = 2
					node.meta &^= (1 << 7)
					node.value = v
				}
			} else {
				if oe {
					rc = 1
					node.value = v
				}
			}
			return
		case -1:
			kvn := newKVNode(k, v, 0, node.prev, node)
			if node.prev == nil {
				n[loc] = kvn
			} else {
				node.prev.next = kvn
			}
			node.prev = kvn
			rc = 0
			return

		}
	}
	n[loc] = newKVNode(k, v, 0, nil, nil)
	rc = 0
	return

}

func isPrime(n uint64) bool {
	if n == 2 || n == 3 {
		return true
	}
	if n <= 1 || n%2 == 0 || n%3 == 0 {
		return false
	}

	i, w := uint64(5), uint64(2)

	for i*i <= n {
		if n%i == 0 {
			return false
		}

		i += w
		w = 6 - w
	}

	return true
}

func generateSipKeys() (key0, key1 uint64, err error) {
	r := make([]byte, 16)
	_, err = rand.Read(r)
	if err != nil {
		return
	}
	key0 = binary.LittleEndian.Uint64(r[:8])
	key1 = binary.LittleEndian.Uint64(r[8:])
	return
}

type kvnode struct {
	key   []byte
	value []byte
	meta  uint8 // metadata field. currently only used for logical delete flag for a record.
	prev  *kvnode
	next  *kvnode
}

// newKVNode returns a new kvnode structure. meta field defaults to 0.
func newKVNode(key, value []byte, meta uint8, prev, next *kvnode) *kvnode {
	return &kvnode{
		key:   key,
		value: value,
		meta:  meta,
		prev:  prev,
		next:  next,
	}
}
