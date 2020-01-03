package dynamic

import (
	"fmt"
)

/*
   question: 对于一个不重复的整数序列，求最长的连续升序子串长度
*/

func longestArr(arr []int) {
	var head []int //子序列第一个值
	var lens []int //子序列长度
	isHead := true
	for _, n := range arr {
		if len(head) == 0 {
			head = append(head, n)
			lens = append(lens, 1)
			continue
		}
		for i, j := range head {
			if n > j {
				lens[i] += 1
				isHead = false
			}
		}
		if isHead {
			head = append(head, n)
			lens = append(lens, 1)
		}
		isHead = true
	}
	fmt.Println(head)
	fmt.Println(lens)
}
