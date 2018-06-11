package hashtable

import (
	"bytes"
	// "log"
	"math/rand"
	"testing"
	"time"
)

type kvpair struct {
	key   []byte
	value []byte
}

type testcase struct {
	name string
	data []kvpair
}

const (
	letterBytes string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%&*() "
)

func TestChainedHashBucketPutAndGet(t *testing.T) {
	cnt := 100000
	bucket := NewChainedHashBucket(uint64(cnt * 3))
	cmap := make(map[string][]byte)
	for i := 0; i < cnt; i++ {
		key := randBytes(20)
		value := randBytes(40)
		bucket.Put(key, value)
		cmap[string(key)] = value
	}
	if uint64(len(cmap)) != bucket.Length() {
		t.Errorf("bucket length %v, is different than map length %v", bucket.Length(), len(cmap))
	}
	for k, v := range cmap {
		if bytes.Compare(v, bucket.Get([]byte(k))) != 0 {
			t.Errorf("key %v has different value in map and hash bucket", k)
		}
	}
}

func TestChainedHashBucketDelete(t *testing.T) {
	cnt := 300000
	bucket := NewChainedHashBucket(uint64(cnt * 3))
	cmap := make(map[string][]byte)
	for i := 0; i < cnt; i++ {
		key := randBytes(10)
		value := randBytes(40)
		bucket.Put(key, value)
		cmap[string(key)] = value
	}
	if uint64(len(cmap)) != bucket.Length() {
		t.Errorf("bucket length %v, is different than map length %v", bucket.Length(), len(cmap))
	}
	// log.Println(bucket.Length())
	for k := range cmap {
		delete(cmap, k)
		bucket.Delete([]byte(k))
		if bucket.Get([]byte(k)) != nil {
			t.Errorf("key %v should have been deleted but is not", k)
		}
		if uint64(len(cmap)) != bucket.Length() {
			t.Errorf("bucket length %v, is different than map length %v", bucket.Length(), len(cmap))
		}
	}
	// log.Println(bucket.Length())
}

func TestChainedHashBucketResizeFull(t *testing.T) {
	cnt := 800000
	bucket := NewChainedHashBucket(uint64(cnt / 8))
	cmap := make(map[string][]byte)
	for i := 0; i < cnt; i++ {
		key := randBytes(10)
		value := randBytes(40)
		bucket.Put(key, value)
		cmap[string(key)] = value
	}
	if uint64(len(cmap)) != bucket.Length() {
		t.Errorf("bucket length %v, is different than map length %v", bucket.Length(), len(cmap))
	}
	// log.Printf("current load factor: %.5f\n", bucket.LoadFactor())
	for f := 4.0; f >= 0.0; f = f - 0.2 {
		err := bucket.ResizeFull(f)
		if err != nil {
			t.Error(err)
		}
		// log.Printf("current load factor: %.5f, bucket capacity %v, f %.2f\n", bucket.LoadFactor(), bucket.Capacity(), f)
		if bucket.LoadFactor() > f && bucket.Capacity() < MaxCapacity {
			t.Errorf("failed to reach target load factor of %.5f, current load factor: %.5f", f, bucket.LoadFactor())
		}
	}
	// depending on system specs this part may cause problems due to memory requirements, especially when all tests are run back to back.
	// err := bucket.ResizeFull(0.0)
	// if err != nil {
	// 	t.Error(err)
	// }
	// if bucket.Capacity() != MaxCapacity {
	// 	t.Errorf("bucket should have reached max capacity")
	// }
}

func TestChainedHashBucketResizePartial(t *testing.T) {
	cnt := 200000
	bucket := NewChainedHashBucket(uint64(cnt / 10))
	cmap := make(map[string][]byte)
	f := 0.3
	for i := 0; i < cnt; i++ {
		key := randBytes(10)
		value := randBytes(40)
		bucket.Put(key, value)
		cmap[string(key)] = value
		var err error
		c := false
		for !c {
			c, err = bucket.ResizePartial(300*time.Nanosecond, f)
			if err != nil {
				t.Error(err)
			}
		}
		if bucket.LoadFactor() > f && bucket.Capacity() < MaxCapacity {
			t.Errorf("failed to reach target load factor of %.5f, current load factor: %.5f", f, bucket.LoadFactor())
		}
	}
	if uint64(len(cmap)) != bucket.Length() {
		t.Errorf("bucket length %v, is different than map length %v", bucket.Length(), len(cmap))
	}

	// depending on system specs this part may cause problems due to memory requirements, especially when all tests are run back to back.
	// var err error
	// c := false
	// for !c {
	// 	c, err = bucket.ResizePartial(300*time.Nanosecond, 0.0)
	// 	if err != nil {
	// 		t.Error(err)
	// 	}
	// }

	// if bucket.Capacity() != MaxCapacity {
	// 	t.Errorf("bucket should have reached max capacity")
	// }
}

func TestChainedHashBucketRandomEvents(t *testing.T) {
	// start with a bucket of up to ~500k items
	cnt := 500000
	bucket := NewChainedHashBucket(uint64(cnt / 2))
	cmap := make(map[string][]byte)
	for i := 0; i < cnt; i++ {
		key := randBytes(15)
		value := randBytes(40)
		bucket.Put(key, value)
		cmap[string(key)] = value
	}

	for i := 0; i < cnt*3; i++ {
		if uint64(len(cmap)) != bucket.Length() {
			t.Errorf("bucket length %v, is different than map length %v", bucket.Length(), len(cmap))
		}
		e := randInt(5)
		switch e {
		case 0:
			// insert 3 elements
			for ix := 0; ix < 3; ix++ {
				key := randBytes(15)
				value := randBytes(40)
				bucket.Put(key, value)
				cmap[string(key)] = value
			}
		case 1:
			// delete an element
			for k := range cmap {
				delete(cmap, k)
				err := bucket.Delete([]byte(k))
				if err != nil {
					t.Error(err)
				}
				break
			}
		case 2:
			// read 5 elements
			c := 0
			for k, v := range cmap {
				if c >= 5 {
					break
				}
				if bytes.Compare(v, bucket.Get([]byte(k))) != 0 {
					t.Errorf("values for key %v are different in hash bucket and map", k)
				}
				c++
			}
		case 3:
			// resize full with target load factor 0.5
			err := bucket.ResizeFull(0.5)
			if err != nil {
				t.Error(err)
			}
			if bucket.LoadFactor() > 0.5 {
				t.Errorf("load factor is above input threshold")
			}
		case 4:
			completed := false
			var err error
			for !completed {
				completed, err = bucket.ResizePartial(500*time.Millisecond, 0.5)
				if err != nil {
					t.Error(err)
				}
			}
			if bucket.LoadFactor() > 0.5 {
				t.Errorf("load factor is above input threshold")
			}
		}
	}
}

func randInt(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100)
}

func randBytes(n int) []byte {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return b
}
