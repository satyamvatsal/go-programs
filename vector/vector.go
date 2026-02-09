package vector

type Vector[T any] struct {
	arr      []T
	Size     int
	Capacity int
}

func NewVector[T any](capacity int) *Vector[T] {
	if capacity <= 0 {
		capacity = 1
	}
	return &Vector[T]{
		arr:      make([]T, capacity),
		Size:     0,
		Capacity: capacity,
	}
}

func (v *Vector[T]) PushBack(data T) {
	if v.Size == v.Capacity {
		newCap := 2*v.Capacity + 1
		newArr := make([]T, newCap)
		copy(newArr, v.arr)
		v.arr = newArr
		v.Capacity = newCap
	}
	v.arr[v.Size] = data
	v.Size++
}

func (v *Vector[T]) PopBack() {
	if v.Size == 0 {
		return
	}
	v.Size--
}
func (v *Vector[T]) Empty() bool {
	return v.Size == 0
}

func (v *Vector[T]) At(index int) (T, bool) {
	var data T
	if index < 0 || index >= v.Size {
		return data, false
	}
	data = v.arr[index]
	return data, true
}

func (v *Vector[T]) Data() []T {
	return v.arr[:v.Size]
}
func (v *Vector[T]) Clear() {
	v.Size = 0
}
