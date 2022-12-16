package ht

import (
	"fmt"
	"strings"

	"github.com/savelyevvv/go-collections/sll"
)

const ErrNoSuchItem = HashTableError("no such item in hash table")

type HashTableError string

func (e HashTableError) Error() string {
	return string(e)
}

type pair[E any] struct {
	key   string
	value E
}

type HashTable[E any] struct {
	buckets []*sll.List[pair[E]]
	size    int
}

func New[E any](capacity int) *HashTable[E] {
	t := &HashTable[E]{
		buckets: make([]*sll.List[pair[E]], capacity),
	}
	for i := range t.buckets {
		t.buckets[i] = sll.New[pair[E]]()
	}
	return t
}

func (t *HashTable[E]) String() string {
	ss := make([]string, 0, t.size)
	for _, bucket := range t.buckets {
		if bucket.Size() > 0 {
			for _, element := range bucket.ToSlice() {
				ss = append(ss, fmt.Sprintf("%q: %v", element.key, element.value))
			}
		}
	}
	return fmt.Sprintf("{%s}", strings.Join(ss, ", "))
}

func (t *HashTable[E]) Size() int {
	return t.size
}

func (t *HashTable[E]) Get(key string) (E, bool) {
	var (
		value E
		ok    bool
	)
	bucket := t.buckets[hash(key, len(t.buckets))]
	if index := bucket.IndexOf(pair[E]{key: key}, func(a, b pair[E]) bool {
		return a.key == b.key
	}); index != -1 {
		value = bucket.Get(index).value
		ok = true
	}
	return value, ok
}

func (t *HashTable[E]) Put(key string, value E) {
	bucket := t.buckets[hash(key, len(t.buckets))]
	p := pair[E]{key: key, value: value}
	if index := bucket.IndexOf(p, func(a, b pair[E]) bool {
		return a.key == b.key
	}); index == -1 {
		bucket.Add(0, p)
		t.size++
	} else {
		bucket.Set(index, p)
	}
}

func (t *HashTable[E]) Delete(key string) {
	bucket := t.buckets[hash(key, len(t.buckets))]
	if ok := bucket.Remove(pair[E]{key: key}, func(a, b pair[E]) bool {
		return a.key == b.key
	}); ok {
		t.size--
	}
}

func hash(key string, size int) int {
	var total int
	for _, r := range key {
		total += int(r)
	}
	return total % size
}
