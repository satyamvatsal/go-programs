package list

type ListNode[T any] struct {
	Data T
	Next *ListNode[T]
}

type List[T any] struct {
	head *ListNode[T]
	size int
	tail *ListNode[T]
}

func NewList[T any]() *List[T] {
	return &List[T]{
		head: nil,
		size: 0,
		tail: nil,
	}
}

func NewListNode[T any]() *ListNode[T] {
	return new(ListNode[T])
}

func (l *List[T]) PushFront(data T) {
	newNode := NewListNode[T]()
	newNode.Data = data
	newNode.Next = l.head
	if l.Empty() {
		l.tail = newNode
	}
	l.head = newNode
	l.size++
}

func (l *List[T]) PushBack(data T) {
	newNode := NewListNode[T]()
	newNode.Data = data
	newNode.Next = nil
	if l.Empty() {
		l.head = newNode
	} else {
		l.tail.Next = newNode
	}
	l.tail = newNode
	l.size++
}

func (l *List[T]) PopFront() {
	if l.Empty() {
		return
	}
	l.head = l.head.Next
}

func (l *List[T]) Empty() bool {
	return l.head == nil
}

func (l *List[T]) Size() int {
	return l.size
}

func (l *List[T]) Front() *ListNode[T] {
	return l.head
}
func (l *List[T]) Back() *ListNode[T] {
	return l.tail
}
