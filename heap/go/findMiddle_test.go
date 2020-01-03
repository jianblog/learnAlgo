package heap

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestFindMiddle(t *testing.T) {
	// 寻找中位数
	rand.Seed(time.Now().UnixNano())
	arr := []int{}
	for i := 0; i < 20; i++ {
		arr = append(arr, rand.Intn(100))
	}

	bigheap := &Heap{[]int{}, -1, 0}
	smallheap := &Heap{[]int{}, -1, 0}
	for _, i := range arr {
		//splitMiddle(bigheap, smallheap, i)
		bigheap.arr[0] = i
		fmt.Println(i)
	}
	fmt.Printf("%v", bigheap)
	fmt.Printf("%v", smallheap)
}
