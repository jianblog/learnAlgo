package heap

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	myHeap := NewHeap(15)
	arr := []int{33, 27, 21, 16, 13, 15, 19, 5, 6, 7, 8, 1, 2, 12}
	for _, i := range arr {
		myHeap.Insert(i)
	}
	fmt.Printf("%v\n", myHeap)

	myHeap.DeleteTop()
	fmt.Printf("%v\n", myHeap)
}
