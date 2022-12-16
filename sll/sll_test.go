package sll

import "fmt"

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
