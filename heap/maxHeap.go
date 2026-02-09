package heap

import "cmp"

type MaxHeap[T cmp.Ordered] struct {
	arr []T
}

func NewMaxHeap[T cmp.Ordered](a []T) MaxHeap[T] {
	data := make([]T, len(a))
	copy(data, a)
	h := MaxHeap[T]{arr: data}
	h.BuildMaxHeap()
	return h
}

func (h *MaxHeap[T]) BuildMaxHeap() {
	n := len(h.arr)
	for i := n/2 - 1; i >= 0; i-- {
		h.Heapify(i)
	}
}

func (h *MaxHeap[T]) Heapify(i int) {
	l, r := h.left(i), h.right(i)
	largest := i
	if l < len(h.arr) && h.arr[largest] < h.arr[l] {
		largest = l
	}
	if r < len(h.arr) && h.arr[largest] < h.arr[r] {
		largest = r
	}
	if largest != i {
		h.arr[largest], h.arr[i] = h.arr[i], h.arr[largest]
		h.Heapify(largest)
	}
}

func (h *MaxHeap[T]) Insert(elem T) {
	h.arr = append(h.arr, elem)
	curr := len(h.arr) - 1
	for curr > 0 && h.arr[curr] > h.arr[h.parent(curr)] {
		p := h.parent(curr)
		h.arr[curr], h.arr[p] = h.arr[p], h.arr[curr]
		curr = p
	}
}

func (h *MaxHeap[T]) ExtractMax() T {
	if len(h.arr) == 0 {
		var zero T
		return zero
	}
	max := h.arr[0]
	lastIdx := len(h.arr) - 1

	h.arr[0] = h.arr[lastIdx]

	var zero T
	h.arr[lastIdx] = zero

	h.arr = h.arr[:lastIdx]

	if len(h.arr) > 0 {
		h.Heapify(0)
	}
	return max
}

func (h *MaxHeap[T]) left(i int) int   { return 2*i + 1 }
func (h *MaxHeap[T]) right(i int) int  { return 2*i + 2 }
func (h *MaxHeap[T]) parent(i int) int { return (i - 1) / 2 }
func (h *MaxHeap[T]) Max() T           { return h.arr[0] }
