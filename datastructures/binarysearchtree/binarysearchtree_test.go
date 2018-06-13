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
		t.Errorf("map and tree lengths are different, map length %v, tree length %v", len(cmap), tree.Length())
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
		t.Errorf("map and tree lengths are different, map length %v, tree length %v", len(cmap), tree.Length())
	}
	for k, v := range cmap {
		tv := tree.Get([]byte(k))
		if bytes.Compare(v, tv) != 0 {
			t.Errorf("map and tree values are different for key %v. map value %v, tree value %v", k, string(v), string(tv))
		}
	}

	for i := 0; i < 100; i++ {
		k := randBytes(1)
		if tree.Get(k) != nil {
			t.Errorf("value for key %v should have been nil", string(k))
		}
	}
}

// TODO: Test minimum
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
		t.Errorf("map and tree lengths are different, map length %v, tree length %v", len(cmap), tree.Length())
	}
	kvmin := tree.Minimum()
	if kvmin == nil {
		t.Error("failed to get minimum")
	}
	found := false
	for k, v := range cmap {
		if bytes.Compare([]byte(k), kvmin.Key) == 0 {
			if bytes.Compare(v, kvmin.Value) != 0 {
				t.Error("minimum returned incorrect value")
			}
			found = true
		}
		if bytes.Compare([]byte(k), kvmin.Key) == -1 {
			t.Error("minimum returned incorrect value")
		}
	}
	if !found {
		t.Error("Minimum returned something irrelevant")
	}
}

// TODO: Test maximum
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
		t.Errorf("map and tree lengths are different, map length %v, tree length %v", len(cmap), tree.Length())
	}
	kvmax := tree.Maximum()
	if kvmax == nil {
		t.Error("failed to get maximum")
	}
	found := false
	for k, v := range cmap {
		if bytes.Compare([]byte(k), kvmax.Key) == 0 {
			if bytes.Compare(v, kvmax.Value) != 0 {
				t.Error("maximum returned incorrect value")
			}
			found = true
		}
		if bytes.Compare([]byte(k), kvmax.Key) == 1 {
			t.Error("maximum returned incorrect value")
		}
	}
	if !found {
		t.Error("Maximum returned something irrelevant")
	}
}

// TODO: Test delete

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
		t.Errorf("map and tree lengths are different, map length %v, tree length %v", len(cmap), tree.Length())
	}

	r := tree.Traverse()

	if len(cmap) != len(r) {
		t.Errorf("map and traversed tree lengths are different, map length %v, tree length %v", len(cmap), len(r))
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
