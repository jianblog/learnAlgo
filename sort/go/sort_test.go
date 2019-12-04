package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	var arr = []int{6, 2, 4, 1, 5, 9}
	ret := BubbleSort(arr[:], 6)
	fmt.Printf("%v", ret)
}

func TestInsertionSort(t *testing.T) {
	var arr = []int{6, 2, 4, 1, 5, 8}
	ret := InsertionSort(arr[0:], 6)
	fmt.Printf("%v", ret)
}

func TestSelectionSort(t *testing.T) {
	var arr = []int{6, 2}
	ret := SelectionSort(arr[:], 2)
	fmt.Printf("%v", ret)
}

func TestMergeSort(t *testing.T) {
	var arr = []int{11, 8, 3, 9, 7, 1, 2, 5}
	ret := MergeSort(arr, 0, 7)
	fmt.Printf("%v", ret)
}
