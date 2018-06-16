package binarysearchtree

import (
	"bytes"
	"math/rand"
	"testing"
	"time"
)

func TestTreeInsert(t *testing.T) {
	cmap := make(map[string][]byte)
	tree := NewTree()

	cnt := 400000

	for i := 0; i < cnt; i++ {
		key := randBytes(6)
		value := randBytes(10)
		cmap[string(key)] = value
		tree.Insert(key, value)
	}

	if uint64(len(cmap)) != tree.Length() {
		t.Fatalf("map and tree lengths are different, map length %v, tree length %v", len(cmap), tree.Length())
	}
}

func TestTreeGet(t *testing.T) {
	cmap := make(map[string][]byte)
	tree := NewTree()

	cnt := 400000

	for i := 0; i < cnt; i++ {
		key := randBytes(6)
		value := randBytes(10)
		cmap[string(key)] = value
		tree.Insert(key, value)
	}

	if uint64(len(cmap)) != tree.Length() {
		t.Fatalf("map and tree lengths are different, map length %v, tree length %v", len(cmap), tree.Length())
	}
	for k, v := range cmap {
		tv := tree.Get([]byte(k))
		if tv == nil {
			t.Fatal("Get did not return any value")
		}
		if bytes.Compare(v, tv) != 0 {
			t.Fatalf("map and tree values are different for key %v. map value %v, tree value %v", k, string(v), string(tv))
		}
	}

	for i := 0; i < 100; i++ {
		k := randBytes(1)
		if tree.Get(k) != nil {
			t.Fatalf("value for key %v should have been nil", string(k))
		}
	}
}

func TestTreeMinimum(t *testing.T) {
	cmap := make(map[string][]byte)
	tree := NewTree()

	cnt := 400000

	for i := 0; i < cnt; i++ {
		key := randBytes(6)
		value := randBytes(10)
		cmap[string(key)] = value
		tree.Insert(key, value)
	}

	if uint64(len(cmap)) != tree.Length() {
		t.Fatalf("map and tree lengths are different, map length %v, tree length %v", len(cmap), tree.Length())
	}
	kvmin := tree.Minimum()
	if kvmin == nil {
		t.Fatal("failed to get minimum")
	}
	found := false
	for k, v := range cmap {
		if bytes.Compare([]byte(k), kvmin.Key) == 0 {
			if bytes.Compare(v, kvmin.Value) != 0 {
				t.Fatal("minimum returned incorrect value")
			}
			found = true
		}
		if bytes.Compare([]byte(k), kvmin.Key) == -1 {
			t.Fatal("minimum returned incorrect value")
		}
	}
	if !found {
		t.Fatal("Minimum returned something irrelevant")
	}
}

func TestTreeMaximum(t *testing.T) {
	cmap := make(map[string][]byte)
	tree := NewTree()

	cnt := 400000

	for i := 0; i < cnt; i++ {
		key := randBytes(6)
		value := randBytes(10)
		cmap[string(key)] = value
		tree.Insert(key, value)
	}

	if uint64(len(cmap)) != tree.Length() {
		t.Fatalf("map and tree lengths are different, map length %v, tree length %v", len(cmap), tree.Length())
	}
	kvmax := tree.Maximum()
	if kvmax == nil {
		t.Fatal("failed to get maximum")
	}
	found := false
	for k, v := range cmap {
		if bytes.Compare([]byte(k), kvmax.Key) == 0 {
			if bytes.Compare(v, kvmax.Value) != 0 {
				t.Fatal("maximum returned incorrect value")
			}
			found = true
		}
		if bytes.Compare([]byte(k), kvmax.Key) == 1 {
			t.Fatal("maximum returned incorrect value")
		}
	}
	if !found {
		t.Fatal("Maximum returned something irrelevant")
	}
}

// TODO: Test delete
func TestTreeDelete(t *testing.T) {

	cmap := make(map[string][]byte)
	cmap2 := make(map[string][]byte)
	tree := NewTree()

	cnt := 400000

	for i := 0; i < cnt; i++ {
		key := randBytes(6)
		value := randBytes(10)
		cmap[string(key)] = value
		cmap2[string(key)] = value
		tree.Insert(key, value)
	}

	if uint64(len(cmap)) != tree.Length() {
		t.Fatalf("map and tree lengths are different, map length %v, tree length %v", len(cmap), tree.Length())
	}
	for k, v := range cmap {
		tv := tree.Get([]byte(k))
		if tv == nil {
			t.Fatal("Get did not return any value")
		}
		if bytes.Compare(v, tv) != 0 {
			t.Fatalf("map and tree values are different for key %v. map value %v, tree value %v", k, string(v), string(tv))
		}
	}
	for k := range cmap2 {
		delete(cmap, k)
		tree.Delete([]byte(k))
		if uint64(len(cmap)) != tree.Length() {
			t.Fatalf("map and tree lengths are different, map length %v, tree length %v", len(cmap), tree.Length())
		}
	}
}

func TestTreeTraverse(t *testing.T) {
	cmap := make(map[string][]byte)
	tree := NewTree()

	cnt := 400000

	for i := 0; i < cnt; i++ {
		key := randBytes(6)
		value := randBytes(10)
		cmap[string(key)] = value
		tree.Insert(key, value)
	}

	if uint64(len(cmap)) != tree.Length() {
		t.Fatalf("map and tree lengths are different, map length %v, tree length %v", len(cmap), tree.Length())
	}

	r := tree.Traverse()

	if len(cmap) != len(r) {
		t.Fatalf("map and traversed tree lengths are different, map length %v, tree length %v", len(cmap), len(r))
	}
}

const (
	letterBytes string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%&*() "
)

func randBytes(n int) []byte {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return b
}
