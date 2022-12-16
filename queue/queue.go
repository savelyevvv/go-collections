package queue

import dlist "github.com/savelyevvv/go-collections/dll"

const ErrEmptyQueue = "queue is empty"

type Queue[E any] struct {
	list *dlist.List[E]
}

func New[E any]() *Queue[E] {
	return &Queue[E]{list: dlist.New[E]()}
}

func (q *Queue[E]) String() string {
	return q.list.String()
}

func (q *Queue[E]) Add(e E) {
	q.list.Append(e)
}

func (q *Queue[E]) Remove() E {
	if q.IsEmpty() {
		panic(ErrEmptyQueue)
	}
	return q.list.Pop(0)
}

func (q *Queue[E]) IsEmpty() bool {
	return q.list.Size() == 0
}

func (q *Queue[E]) Element() E {
	if q.IsEmpty() {
		panic(ErrEmptyQueue)
	}
	return q.list.Get(0)
}
