package queue

import (
	"github.com/satyamvatsal/go-programs/list"
)

type Queue[T any] struct {
	l *list.List[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		l: list.NewList[T](),
	}
}

func (q *Queue[T]) Enqueue(data T) {
	q.l.PushBack(data)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	return q.l.PopFront()
}

func (q *Queue[T]) Front() (data T, ok bool) {
	if q.l.Empty() {
		return data, false
	}
	return q.l.Front().Data, true
}

func (q *Queue[T]) Empty() bool {
	return q.l.Empty()
}
