package heap

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	myHeap := NewHeap(15)
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//arr := []int{33, 27, 21, 16, 13, 15, 19, 5, 6, 7, 8, 1, 2, 12}
	for _, i := range arr {
		myHeap.Insert(i)
	}
	fmt.Printf("%v\n", *myHeap)

	myHeap.DeleteTop()
	fmt.Printf("删除顶部元素:\n %v\n", *myHeap)
}

func TestHeapify(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	Heapify(arr)
	fmt.Printf("数组heapify:\n%v", arr)
}
