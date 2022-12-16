package sll

import (
	"fmt"
)

func ExampleList_IndexOf_emptyList() {
	list := New[int]()
	index := list.IndexOf(10, func(a, b int) bool { return a == b })
	fmt.Println(index)
	// Output: -1
}

func ExampleList_IndexOf_hasElement() {
	list := New[int]()
	list.Append(10)
	list.Append(15)
	list.Append(20)
	index := list.IndexOf(15, func(a, b int) bool { return a == b })
	fmt.Println(index)
	// Output: 1
}

func ExampleList_IndexOf_doesNotHaveElement() {
	list := New[int]()
	list.Append(10)
	list.Append(15)
	list.Append(20)
	index := list.IndexOf(25, func(a, b int) bool { return a == b })
	fmt.Println(index)
	// Output: -1
}

func ExampleList_ToSlice_emptyList() {
	list := New[int]()
	slice := list.ToSlice()
	fmt.Println(slice)
	// Output: []
}

func ExampleList_ToSlice_nonEmptyList() {
	list := New[int]()
	list.Append(10)
	list.Append(15)
	list.Append(20)
	slice := list.ToSlice()
	fmt.Println(slice)
	// Output: [10 15 20]
}

func ExampleList_Remove_contains() {
	list := New[int]()
	list.Append(10)
	list.Append(15)
	list.Append(20)
	rsl := list.Remove(func(e int) bool { return e == 15 })
	fmt.Println(rsl, list)
	// Output: true [10, 20]
}

func ExampleList_Remove_doesNotContain() {
	list := New[int]()
	list.Append(10)
	list.Append(15)
	list.Append(20)
	rsl := list.Remove(func(e int) bool { return e == 25 })
	fmt.Println(rsl, list)
	// Output: false [10, 15, 20]
}

func ExampleList_Find_contains() {
	list := New[int]()
	list.Append(10)
	list.Append(15)
	list.Append(20)
	element, ok := list.Find(func(e int) bool { return e%2 == 1 })
	fmt.Println(element, ok)
	// Output: 15 true
}

func ExampleList_Find_doesNotContain() {
	list := New[int]()
	list.Append(10)
	list.Append(15)
	list.Append(20)
	element, ok := list.Find(func(e int) bool { return e == 25 })
	fmt.Println(element, ok)
	// Output: 0 false
}
