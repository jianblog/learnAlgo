package kmp

import "fmt"

func ViolentMatch(s string, p string) {
	//暴力法进行字符匹配
	lenS := len(s)
	lenP := len(p)
	fmt.Println("length: ", lenS, lenP)
	i, j := 0, 0
	for i < lenS && j < lenP {
		fmt.Println(i, j)
		if s[i] == p[j] {
			i += 1
			j += 1
		} else {
			i = i - (j - 1)
			j = 0
		}
	}
	if j == lenP {
		fmt.Println(i - lenP)
	} else {
		fmt.Println("no found")
	}
}
