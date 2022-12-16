package ht

import (
	"fmt"
	"strings"

	"github.com/savelyevvv/go-collections/sll"
)

// HashTable
// .Insert
// .Search
// .Delete

// Hash()

type HashTable[E any] struct {
	buckets []*sll.List[E]
	size    int
}

func New[E any](size int) *HashTable[E] {
	t := &HashTable[E]{
		buckets: make([]*sll.List[E], size),
	}
	for i := range t.buckets {
		t.buckets[i] = sll.New[E]()
	}
	return t
}

func (t *HashTable[E]) String() string {
	ss := make([]string, 0, t.size)
	for _, bucket := range t.buckets {
		if bucket.Size() > 0 {
			for _, element := range bucket.ToSlice() {
				ss = append(ss, fmt.Sprintf("%v", element))
			}
		}
	}
	return fmt.Sprintf("[%s]", strings.Join(ss, ","))
}
