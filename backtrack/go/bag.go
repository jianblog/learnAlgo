package bag

/*
    背包问题: 背包总承重为W公斤, n个物品。
	期望选择几件物品，在不超过背包承重前提下，让背包中物品总重最大
*/

import (
	"fmt"
)

// 背包总重
var maxW int = 0

//items := [8]{23, 12, 2, 56, 121, 33, 5, 1}  物品重量

func putBag(i int, allWeight int, items []int, capacity int) {
	if allWeight == capacity || i == 9 {
		if allWeight > maxW {
			maxW = allWeight
			return
		}
		putbag(i+1, allWeight, items, capacity)
		if allWeight+items[i] <= capacity {
			putbag(i+1, allWeight+items[i], items, capacity)
		}
	}
}
