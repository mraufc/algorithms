package linkedlist

import (
	"bytes"
	"sync"
)

// LinkedList is a thread-safe, ordered, circular doubly linked list
type LinkedList struct {
	mtx    *sync.RWMutex
	head   *node
	length uint64 // number distinct keys in the list
	count  uint64 // number of total key - value pairs in the list
}

// NewLinkedList returns a new linked list with zero values for head and lenght
func NewLinkedList() *LinkedList {
	return &LinkedList{mtx: new(sync.RWMutex)}
}

type node struct {
	key        []byte
	values     [][]byte
	next, prev *node
}

func newNode(key, value []byte) *node {
	n := new(node)
	n.key = key
	n.values = make([][]byte, 1)
	n.values[0] = value
	n.prev, n.next = n, n
	return n
}

// Min returns the minimum key and its values
func (l *LinkedList) Min() (key []byte, values [][]byte) {
	if l.length == 0 {
		return
	}
	n := l.head
	key, values = n.keyContents(), n.valueContents()
	return
}

// Max returns the maximum key and its values
func (l *LinkedList) Max() (key []byte, values [][]byte) {
	if l.length == 0 {
		return
	}
	n := l.head.prev
	key, values = n.keyContents(), n.valueContents()
	return
}

func (n *node) keyContents() []byte {
	// need to return copies of key and values or else any modification to returned values will modify the linked list
	key := make([]byte, len(n.key))
	copy(key, n.key)
	return key
}

func (n *node) valueContents() [][]byte {
	values := make([][]byte, len(n.values))
	for i := range values {
		x := make([]byte, len(n.values[i]))
		copy(x, n.values[i])
		values[i] = x
	}
	return values
}

// Delete deletes a key and associated values from the list
func (l *LinkedList) Delete(key []byte) {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	n := l.head
	// check minimum and maximum keys in the list
	if l.length > 0 {
		if bytes.Compare(key, n.key) == -1 || bytes.Compare(key, n.prev.key) == 1 {
			return
		}
	}

	for i := uint64(0); i < l.length; i++ {
		if bytes.Compare(key, n.key) == -1 {
			break
		}
		if bytes.Compare(key, n.key) == 0 {
			l.length--
			l.count -= uint64(len(n.values))
			if l.length == 0 {
				l.head = nil
				break
			}
			n.next.prev = n.prev
			n.prev.next = n.next
			if i == uint64(0) {
				l.head = l.head.next
			}
			break
		}
		n = n.next
	}
	return

}

// Search returns the set of values for a given key.
// Returns nil when there are no keys match the input.
func (l *LinkedList) Search(key []byte) [][]byte {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	n := l.head
	// check minimum and maximum keys in the list
	if l.length > 0 {
		if bytes.Compare(key, n.key) == -1 || bytes.Compare(key, n.prev.key) == 1 {
			return nil
		}
	}

	for i := uint64(0); i < l.length; i++ {
		if bytes.Compare(key, n.key) == -1 {
			break
		}
		if bytes.Compare(key, n.key) == 0 {
			return n.valueContents()
		}
		n = n.next
	}
	return nil
}

// Insert inserts a key value pair to the linked list
func (l *LinkedList) Insert(key, value []byte) {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	if l.length == 0 {
		l.length++
		l.count++
		l.head = newNode(key, value)
		return
	}
	// handle "special" cases where key will be the least or largest element in the list
	if bytes.Compare(key, l.head.key) == -1 {
		l.length++
		l.count++
		n := newNode(key, value)
		l.head.addAsPrev(n)
		l.head = n
		return
	}
	if bytes.Compare(key, l.head.prev.key) == 1 {
		l.length++
		l.count++
		l.head.addAsPrev(newNode(key, value))
		return
	}
	n := l.head
	for i := uint64(0); i < l.length; i++ {

		if bytes.Compare(key, n.key) == 0 {
			l.count++
			n.values = append(n.values, value)
			break
		}
		if bytes.Compare(key, n.key) == -1 {
			l.count++
			l.length++
			n.addAsPrev(newNode(key, value))
		}

		n = n.next

	}
}

// Length returns number of distinct keys in the list
func (l *LinkedList) Length() uint64 {
	return l.length
}

// Count returns number of key value pairs in the list
func (l *LinkedList) Count() uint64 {
	return l.count
}

func (n *node) addAsPrev(nn *node) {
	nn.next = n
	nn.prev = n.prev
	n.prev.next = nn
	n.prev = nn
}
