package datastructures

import (
	"bytes"
	"testing"
)

type kvpair struct {
	key    []byte
	values [][]byte
}

type testcase struct {
	name string
	data []kvpair
}

func TestLinkedListInsert(t *testing.T) {
	type expected struct {
		length uint64
		count  uint64
	}
	ex := []expected{
		expected{
			uint64(1),
			uint64(2),
		},
		expected{
			uint64(3),
			uint64(6),
		},
		expected{
			uint64(6),
			uint64(9),
		},
	}
	tests := testData()
	for ix, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			l := NewLinkedList()
			for _, kv := range tc.data {
				for _, v := range kv.values {
					l.Insert(kv.key, v)
				}
			}
			if l.Length() != ex[ix].length || l.Count() != ex[ix].count {
				t.FailNow()
			}

		})
	}
}

func TestLinkedListSearch(t *testing.T) {
	tests := testData()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			l := NewLinkedList()
			for _, kv := range tc.data {
				for _, v := range kv.values {
					l.Insert(kv.key, v)
				}
			}
			for _, kv := range tc.data {
				r := l.Search(kv.key)
				t.Log(len(r))
				for i, v := range kv.values {
					if bytes.Compare(r[i], v) != 0 {
						t.FailNow()
					}
				}
				if l.Search([]byte("key0")) != nil {
					t.FailNow()
				}
				if l.Search([]byte("key9")) != nil {
					t.FailNow()
				}
				if l.Search([]byte("key6")) != nil {
					t.FailNow()
				}
			}
		})
	}
}

func TestLinkedListMinAndMax(t *testing.T) {
	tests := testData()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			l := NewLinkedList()
			for _, kv := range tc.data {
				for _, v := range kv.values {
					l.Insert(kv.key, v)
				}
			}
			minkey, minvalues := l.Min()
			maxkey, maxvalues := l.Max()
			for _, kv := range tc.data {
				if bytes.Compare(kv.key, minkey) == -1 {
					t.FailNow()
				}
				if bytes.Compare(kv.key, maxkey) == 1 {
					t.FailNow()
				}
				if bytes.Compare(kv.key, minkey) == 0 {
					for i, v := range kv.values {
						if bytes.Compare(minvalues[i], v) != 0 {
							t.FailNow()
						}
					}
				}
				if bytes.Compare(kv.key, maxkey) == 0 {
					for i, v := range kv.values {
						if bytes.Compare(maxvalues[i], v) != 0 {
							t.FailNow()
						}
					}
				}
			}
		})
	}
}

func TestLinkedListDelete(t *testing.T) {
	tests := testData()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			l := NewLinkedList()
			for _, kv := range tc.data {
				for _, v := range kv.values {
					l.Insert(kv.key, v)
				}
			}
			for i := len(tc.data) - 1; i >= 0; i-- {
				if l.Search(tc.data[i].key) == nil {
					t.FailNow()
				}
				l.Delete(tc.data[i].key)
				if l.Search(tc.data[i].key) != nil {
					t.FailNow()
				}
			}
		})
	}
}

func testData() []testcase {
	return []testcase{
		{
			"List_1",
			[]kvpair{
				kvpair{
					[]byte("key1"),
					[][]byte{
						[]byte("value1_1"),
						[]byte("value1_2"),
					},
				},
			},
		},
		{
			"List_2",
			[]kvpair{
				kvpair{
					[]byte("key2"),
					[][]byte{
						[]byte("value2_1"),
						[]byte("value2_2"),
					},
				},
				kvpair{
					[]byte("key5"),
					[][]byte{
						[]byte("value5_1"),
						[]byte("value5_2"),
						[]byte("value5_3"),
					},
				},
				kvpair{
					[]byte("key3"),
					[][]byte{
						[]byte("value3_1"),
					},
				},
			},
		},
		{
			"List_3",
			[]kvpair{
				kvpair{
					[]byte("key2"),
					[][]byte{
						[]byte("value2_1"),
						[]byte("value2_2"),
					},
				},
				kvpair{
					[]byte("key5"),
					[][]byte{
						[]byte("value5_1"),
						[]byte("value5_2"),
						[]byte("value5_3"),
					},
				},
				kvpair{
					[]byte("key3"),
					[][]byte{
						[]byte("value3_1"),
					},
				},
				kvpair{
					[]byte("key1"),
					[][]byte{
						[]byte("value1_1"),
					},
				},
				kvpair{
					[]byte("key4"),
					[][]byte{
						[]byte("value4_1"),
					},
				},
				kvpair{
					[]byte("key7"),
					[][]byte{
						[]byte("value7_1"),
					},
				},
			},
		},
	}
}
