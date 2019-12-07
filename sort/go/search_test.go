package sort

import (
	"fmt"
	"testing"
)

func TestBSearch(t *testing.T) {
	arr := []int{2, 4, 5, 9, 14, 23, 34, 58, 79, 92}
	ret := bSearch(arr, 14)
	fmt.Printf("%v, index=%d", arr, ret)
}
func TestOtherSearch1(t *testing.T) {
	arr1 := []int{2, 4, 14, 14, 14, 56, 77, 81, 92, 101}
	ret := binaryOtherSearch1(arr1, 14)
	fmt.Printf("%v, index=%d\n", arr1, ret)

	arr2 := []int{2, 2, 2, 2, 2, 2, 2, 2, 2}
	ret2 := binaryOtherSearch1(arr2, 2)
	fmt.Printf("%v, index=%d\n", arr2, ret2)

}

func TestOtherSearch3(t *testing.T) {
	arr1 := []int{2, 4, 14, 14, 25, 56, 77, 81, 92}
	v := []int{1, 2, 14, 92, 95}
	for _, i := range v {
		ret := binaryOtherSearch4(arr1, i)
		fmt.Printf("%v, %d index=%d\n", arr1, i, ret)
	}
}
