package ht

import "github.com/savelyevvv/go-collections/sll"

// HashTable
// .Insert
// .Search
// .Delete

// Hash()

type HashTable[E any] struct {
	buckets []sll.List[E]
	size    int
}

func (t *HashTable[E]) String() string {
	for _, bucket := range t.buckets {
		if bucket.Size() > 0 {

		}
	}
}
