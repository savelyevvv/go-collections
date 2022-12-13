package dlist

import (
	"fmt"
	"strings"
)

const ErrEmptyList = "list is empty"
const ErrIndexOutOfBounds = "index out of bounds"

type Cell[E any] struct {
	next *Cell[E]
	prev *Cell[E]
	data E
}

func (n *Cell[E]) String() string {
	return fmt.Sprintf("%v", n.data)
}

type List[E any] struct {
	head *Cell[E]
	tail *Cell[E]
	size int
}

func NewList[E any]() *List[E] {
	tail := &Cell[E]{}
	head := &Cell[E]{}
	head.next = tail
	tail.prev = head
	return &List[E]{head: head, tail: tail}
}

func (l *List[E]) String() string {
	ss := make([]string, 0, l.size)
	temp := l.head.next
	for temp.next != nil {
		ss = append(ss, fmt.Sprintf("%v", temp.data))
		temp = temp.next
	}
	return fmt.Sprintf("[%s]", strings.Join(ss, ", "))
}

func (l *List[E]) Size() int {
	return l.size
}

func (l *List[E]) Get(index int) E {
	if index < 0 || index >= l.size {
		panic(ErrIndexOutOfBounds)
	}
	return l.getCell(index).data
}

func (l *List[E]) Append(e E) {
	l.Add(l.size, e)
}

func (l *List[E]) Add(index int, e E) {
	if index < 0 || index > l.size {
		panic(ErrIndexOutOfBounds)
	}
	prev := l.getCell(index - 1)
	insertCell(prev, &Cell[E]{data: e})
	l.size++
}

func (l *List[E]) Pop(index int) E {
	if index < 0 || index >= l.size {
		panic(ErrIndexOutOfBounds)
	}
	prev := l.getCell(index - 1)
	data := prev.next.data
	deleteCell(prev)
	l.size--
	return data
}

func (l *List[E]) getCell(index int) *Cell[E] {
	var cell *Cell[E]
	if index < l.size/2 {
		cell = l.getCellFromHead(index)
	} else {
		cell = l.getCellFromTail(index)
	}
	return cell
}

func (l *List[E]) getCellFromHead(index int) *Cell[E] {
	var (
		position = -1
		cell     = l.head
	)
	for position < index {
		position++
		cell = cell.next
	}
	return cell
}

func (l *List[E]) getCellFromTail(index int) *Cell[E] {
	var (
		position = l.size
		cell     = l.tail
	)
	for position > index {
		position--
		cell = cell.prev
	}
	return cell
}

func insertCell[E any](prev, cell *Cell[E]) {
	cell.next = prev.next
	prev.next = cell
	cell.next.prev = cell
	cell.prev = prev
}

func deleteCell[E any](prev *Cell[E]) {
	prev.next.next.prev = prev
	prev.next = prev.next.next
}
