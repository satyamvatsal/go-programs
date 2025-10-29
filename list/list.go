package list

type ListNode[T any] struct {
	Data T
	Next *ListNode[T]
}

type List[T any] struct {
	head *ListNode[T]
	tail *ListNode[T]
	size int
}

func NewList[T any]() *List[T] {
	return &List[T]{}
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

func (l *List[T]) PopFront() (data T, ok bool) {
	if l.Empty() {
		return data, false
	}
	data = l.head.Data
	l.head = l.head.Next
	if l.head == nil {
		l.tail = nil
	}
	l.size--
	return data, true
}

func (l *List[T]) PopBack() (data T, ok bool) {
	if l.Empty() {
		return data, false
	}
	head := l.head
	data = l.tail.Data
	if head.Next == nil {
		// contains single element
		l.head = nil
		l.tail = nil
		l.size--
		return data, true
	}
	for head.Next != l.tail {
		head = head.Next
	}
	head.Next = nil
	l.tail = head
	l.size--
	return data, true
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
