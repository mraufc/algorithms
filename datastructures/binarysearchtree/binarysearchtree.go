package binarysearchtree

import (
	"bytes"
	"sync"
)

// Tree is a simple thread-safe binary search tree implementation
type Tree struct {
	mtx    *sync.RWMutex
	root   *node
	length uint64
}

type node struct {
	parent, left, right *node
	key, value          []byte
}

// KVPair is a simple key - value pair structure
type KVPair struct {
	Key, Value []byte
}

func newNode(key, value []byte) *node {
	return &node{key: key, value: value}
}

// NewTree returns a new empty Tree with its root nil
func NewTree() *Tree {
	return &Tree{mtx: new(sync.RWMutex)}
}

// Length returns the number of elements in  the tree
func (t *Tree) Length() uint64 {
	t.mtx.RLock()
	r := t.length
	t.mtx.RUnlock()
	return r
}

func (n *node) locate(k []byte) *node {
	switch bytes.Compare(k, n.key) {
	case 0:
		return n
	case 1:
		if n.right != nil {
			return n.right.locate(k)
		}
	case -1:
		if n.left != nil {
			return n.left.locate(k)
		}
	}
	return nil
}

// Maximum returns key value pair for maximum key in the tree
func (t *Tree) Maximum() *KVPair {
	var n *node
	t.mtx.RLock()
	if t.root != nil {
		n = t.root.maximum()
	}
	t.mtx.RUnlock()
	if n == nil {
		return nil
	}
	return &KVPair{Key: n.key, Value: n.value}
}

// Minimum returns key value pair for minimum key in the tree
func (t *Tree) Minimum() *KVPair {
	var n *node
	t.mtx.RLock()
	if t.root != nil {
		n = t.root.minimum()
	}
	t.mtx.RUnlock()
	if n == nil {
		return nil
	}
	return &KVPair{Key: n.key, Value: n.value}
}

// minimum returns the minimum key value pair given a node
func (n *node) minimum() *node {
	v := n
	for v.left != nil {
		v = v.left
	}
	return v
}

// maximum returns the minimum key value pair given a node
func (n *node) maximum() *node {
	v := n
	for v.right != nil {
		v = v.right
	}
	return v
}

// Delete removes the node corresponding to the input key from the tree
func (t *Tree) Delete(key []byte) {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	if t.root == nil {
		return
	}
	n := t.root.locate(key)
	if n == nil {
		return
	}
	if n.left == nil {
		t.transplant(n, n.right) // n has no left child
	} else if n.right == nil {
		t.transplant(n, n.left) // n has a left child but no right child
	} else {
		y := n.right.minimum()
		if y.parent != n {
			t.transplant(y, y.right)
			y.right = n.right
			y.right.parent = y
		}
		t.transplant(n, y)
		y.left = n.left
		y.left.parent = y
	}
	t.length--
}

// transplant replaces subtree that's rooted at u with the subtree that's rooted at v
func (t *Tree) transplant(u, v *node) {
	if u.parent == nil {
		t.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	if v != nil {
		v.parent = u.parent
	}
}

// Get returns the value corresponding to an input key. Returns nil if the key does not exist.
func (t *Tree) Get(key []byte) []byte {
	t.mtx.RLock()
	var v []byte
	if t.root != nil {
		n := t.root.locate(key)
		if n != nil {
			v = n.value
		}
	}
	t.mtx.RUnlock()
	return v
}

// Traverse is inorder tree walk implementation
func (t *Tree) Traverse() []*KVPair {
	t.mtx.RLock()
	if t.root == nil {
		return nil
	}
	r := t.root.traverse()
	t.mtx.RUnlock()
	return r
}

func (n *node) traverse() []*KVPair {
	var r, l, result []*KVPair
	if n.left != nil {
		l = n.left.traverse()
	}
	result = make([]*KVPair, 1)
	result[0] = &KVPair{Key: n.key, Value: n.value}
	if n.right != nil {
		r = n.right.traverse()
	}
	return append(append(l, result...), r...)
}

// Insert inserts a key value pair to the Tree. If key already exists, value is updated.
func (t *Tree) Insert(key, value []byte) {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	if t.root == nil {
		t.root = newNode(key, value)
		t.length = 1
		return
	}
	r := t.root
	var y *node
	for r != nil {
		y = r
		switch bytes.Compare(r.key, key) {
		case 0:
			r.value = value
			return
		case 1:
			r = r.left
		case -1:
			r = r.right
		}
	}
	z := newNode(key, value)
	z.parent = y
	switch bytes.Compare(y.key, key) {
	case 1:
		y.left = z
	case -1:
		y.right = z
	}
	t.length++
}
