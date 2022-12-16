package ht

import (
	"fmt"
)

func ExampleNew() {
	capacity := 100
	table := New[int](capacity)
	fmt.Println(table)
	// Output: {}
}

func ExampleHashTable_Put() {
	capacity := 100
	table := New[int](capacity)
	table.Put("lora", 25)
	table.Put("alex", 35)
	table.Put("xela", 45)
	fmt.Println(table, table.Size())
	// Output: {"xela": 45, "alex": 35, "lora": 25} 3
}

func ExampleHashTable_Put_replace() {
	capacity := 100
	table := New[int](capacity)
	table.Put("lora", 25)
	table.Put("alex", 35)
	table.Put("xela", 45)
	table.Put("alex", 55)
	fmt.Println(table, table.Size())
	// Output: {"alex": 55, "xela": 45, "lora": 25} 3
}

func ExampleHashTable_Get_contains() {
	capacity := 100
	table := New[int](capacity)
	table.Put("lora", 25)
	table.Put("alex", 35)
	table.Put("xela", 45)
	fmt.Println(table.Get("alex"))
	// Output: 35 true
}

func ExampleHashTable_Get_doesNotContain() {
	size := 100
	table := New[int](size)
	table.Put("lora", 25)
	table.Put("alex", 35)
	table.Put("xela", 45)
	fmt.Println(table.Get("lexa"))
	// Output: 0 false
}

func ExampleHashTable_Delete_contains() {
	capacity := 100
	table := New[int](capacity)
	table.Put("lora", 25)
	table.Put("alex", 35)
	table.Put("xela", 45)
	table.Delete("alex")
	fmt.Println(table, table.Size())
	// Output: {"xela": 45, "lora": 25} 2
}

func ExampleHashTable_Delete_doesNotContain() {
	capacity := 100
	table := New[int](capacity)
	table.Put("lora", 25)
	table.Put("alex", 35)
	table.Put("xela", 45)
	table.Delete("lexa")
	fmt.Println(table, table.Size())
	// Output: {"xela": 45, "alex": 35, "lora": 25} 3
}
