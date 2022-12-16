package ht

import "fmt"

func ExampleNew() {
	size := 5
	table := New[int](size)
	fmt.Println(table)
	// Output: []
}
