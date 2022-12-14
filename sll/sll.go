package sll

import (
	"fmt"
	"strings"
)

const ErrEmptyList = "list is empty"
const ErrIndexOutOfBounds = "index out of bounds"

type Cell[E any] struct {
	data E
	next *Cell[E]
}

func (n *Cell[E]) String() string {
	return fmt.Sprintf("%v", n.data)
}

type List[E any] struct {
	head *Cell[E]
	size int
}

func New[E any]() *List[E] {
	var limiter Cell[E]
	return &List[E]{head: &limiter}
}

func (l *List[E]) String() string {
	ss := make([]string, 0, l.size)
	temp := l.head.next
	for temp != nil {
		ss = append(ss, fmt.Sprintf("%v", temp.data))
		temp = temp.next
	}
	return fmt.Sprintf("[%s]", strings.Join(ss, ", "))
}

func (l *List[E]) Size() int {
	return l.size
}

func (l *List[E]) Append(e E) {
	temp := l.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = &Cell[E]{data: e}
	l.size++
}

func (l *List[E]) Add(index int, e E) {
	if index < 0 || index > l.size {
		panic(ErrIndexOutOfBounds)
	}
	cell := l.getCell(index - 1)
	cell.next = &Cell[E]{data: e, next: cell.next}
	l.size++
}

func (l *List[E]) Pop(index int) E {
	if index < 0 || index >= l.size {
		panic(ErrIndexOutOfBounds)
	}
	cell := l.getCell(index - 1)
	data := cell.next.data
	cell.next = cell.next.next
	l.size--
	return data
}

func (l *List[E]) Get(index int) E {
	if index < 0 || index >= l.size {
		panic(ErrIndexOutOfBounds)
	}
	return l.getCell(index).data
}

func (l *List[E]) Set(index int, element E) {
	if index < 0 || index >= l.size {
		panic(ErrIndexOutOfBounds)
	}
	l.getCell(index).data = element
}

func (l *List[E]) FindIndex(callback func(e E) bool) int {
	var (
		index int
		temp  *Cell[E]
	)
	temp = l.head.next
	for temp != nil {
		if callback(temp.data) {
			break
		}
		index++
		temp = temp.next
	}
	if index == l.size {
		index = -1
	}
	return index
}

func (l *List[E]) Remove(callback func(e E) bool) bool {
	var result bool
	node := l.findNodeBefore(callback)
	if node != nil {
		deleteAfter(node)
		result = true
	}
	return result
}

func (l *List[E]) Find(callback func(e E) bool) (E, bool) {
	var (
		element E
		ok      bool
	)
	temp := l.head.next
	for temp != nil {
		if callback(temp.data) {
			element = temp.data
			ok = true
			break
		}
		temp = temp.next
	}
	return element, ok
}

func (l *List[E]) ToSlice() []E {
	var (
		slice []E
		temp  *Cell[E]
	)
	temp = l.head.next
	for temp != nil {
		slice = append(slice, temp.data)
		temp = temp.next
	}
	return slice
}

func (l *List[E]) getCell(index int) *Cell[E] {
	var (
		position = -1
		temp     = l.head
	)
	for position < index {
		temp = temp.next
		position++
	}
	return temp
}

func (l *List[E]) findNode(callback func(e E) bool) *Cell[E] {
	var (
		node *Cell[E]
		temp *Cell[E]
	)
	temp = l.head.next
	for temp != nil {
		if callback(temp.data) {
			node = temp
			break
		}
		temp = temp.next
	}
	return node
}

func (l *List[E]) findNodeBefore(callback func(e E) bool) *Cell[E] {
	var (
		node *Cell[E]
		temp *Cell[E]
	)
	temp = l.head
	for temp.next != nil {
		if callback(temp.next.data) {
			node = temp
			break
		}
		temp = temp.next
	}
	return node
}

func deleteAfter[E any](previous *Cell[E]) {
	previous.next = previous.next.next
}

func insertAfter[E any](previous, node *Cell[E]) {
	node.next = previous.next
	previous.next = node
}
