package kmp

/*
  next函数值计算
*/
import (
	"fmt"
)

func getNexts(s string) {
	lenS := len(s)
	arrNext := make([]int, lenS)
	arrNext[0] = -1

	k := -1
	for i := 1; i < lenS; i++ {
		for k != -1 && s[k+1] != s[i] {
			k = arrNext[k]
		}
		if s[k+1] == s[i] {
			k += 1
		}
		arrNext[i] = k
	}
	fmt.Println(arrNext)
}
