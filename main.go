package main

import (
	"fmt"

	"github.com/satyamvatsal/go-programs/vector"
)

func main() {
	v := vector.NewVector[int](1)
	for i := range 10 {
		v.PushBack(i)
		fmt.Println(v.Size)
	}
	for _ = range 4 {
		v.PopBack()
	}
	fmt.Printf("%+v\n", v)
}
