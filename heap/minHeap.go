package heap

import "cmp"

type Heap[T cmp.Ordered] struct {
	arr []T
}

func NewHeap[T cmp.Ordered](a []T) Heap[T] {
	data := make([]T, len(a))
	copy(data, a)
	h := Heap[T]{arr: data}
	h.BuildMinHeap()
	return h
}

func (h *Heap[T]) BuildMinHeap() {
	n := len(h.arr)
	for i := n/2 - 1; i >= 0; i-- {
		h.Heapify(i)
	}
}

func (h *Heap[T]) Heapify(i int) {
	l, r := h.left(i), h.right(i)
	smallest := i
	if l < len(h.arr) && h.arr[smallest] > h.arr[l] {
		smallest = l
	}
	if r < len(h.arr) && h.arr[smallest] > h.arr[r] {
		smallest = r
	}
	if smallest != i {
		h.arr[smallest], h.arr[i] = h.arr[i], h.arr[smallest]
		h.Heapify(smallest)
	}
}

func (h *Heap[T]) Insert(elem T) {
	h.arr = append(h.arr, elem)
	curr := len(h.arr) - 1
	for curr > 0 && h.arr[curr] < h.arr[h.parent(curr)] {
		p := h.parent(curr)
		h.arr[curr], h.arr[p] = h.arr[p], h.arr[curr]
		curr = p
	}
}

func (h *Heap[T]) ExtractMin() T {
	if len(h.arr) == 0 {
		var zero T
		return zero
	}
	min := h.arr[0]
	lastIdx := len(h.arr) - 1

	h.arr[0] = h.arr[lastIdx]

	var zero T
	h.arr[lastIdx] = zero

	h.arr = h.arr[:lastIdx]

	if len(h.arr) > 0 {
		h.Heapify(0)
	}
	return min
}

func (h *Heap[T]) left(i int) int   { return 2*i + 1 }
func (h *Heap[T]) right(i int) int  { return 2*i + 2 }
func (h *Heap[T]) parent(i int) int { return (i - 1) / 2 }
func (h *Heap[T]) Min() T           { return h.arr[0] }
