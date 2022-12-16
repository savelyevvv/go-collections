package stack

import "github.com/savelyevvv/go-collections/sll"

const ErrEmptyStack = "stack is empty"

type Stack[E any] struct {
	list *sll.List[E]
}

func New[E any]() *Stack[E] {
	return &Stack[E]{list: sll.New[E]()}
}

func (s *Stack[E]) String() string {
	return s.list.String()
}

func (s *Stack[E]) IsEmpty() bool {
	return s.list.Size() == 0
}

func (s *Stack[E]) Peek() E {
	if s.IsEmpty() {
		panic(ErrEmptyStack)
	}
	return s.list.Get(0)
}

func (s *Stack[E]) Push(item E) E {
	s.list.Add(0, item)
	return item
}

func (s *Stack[E]) Pop() E {
	if s.IsEmpty() {
		panic(ErrEmptyStack)
	}
	return s.list.Pop(0)
}
