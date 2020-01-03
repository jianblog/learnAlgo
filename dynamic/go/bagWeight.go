package dynamic

import (
	"fmt"
)

func calState() {
	items := []int{23, 12, 2, 56, 121, 33, 5, 1}
	capacity := 100
	var states [8][101]bool
	states[0][0] = true

	if items[0] <= capacity {
		states[0][items[0]] = true
	}

	for i := 1; i < 8; i++ {
		for j := 0; j < 100; j++ {
			// 不放[i]物品,总重与之前相同
			if states[i-1][j] {
				states[i][j] = states[i-1][j]
			}
			// 放[i]物品，总重加当前i物品,但是之后总重不能超过承重
			if j+items[i] <= capacity {
				if states[i-1][j] {
					states[i][j+items[i]] = states[i-1][j]
				}
			}
		}
	}
	for j := 100; j >= 0; j-- {
		if states[7][j] {
			fmt.Println(j)
			return
		}
	}
}

func calValue() {
	// 在满足不超重前提下，要求价值最大
	items := [5]int{2, 2, 4, 6, 3}
	values := [5]int{3, 4, 8, 9, 6}
	w := 9

	var states [5][10]int
	// 由于当前数组值中0具有计算意义，需要同不存在的状态为-1区分
	for i := 0; i < 5; i++ {
		for j := 0; j < 10; j++ {
			states[i][j] = -1
		}
	}
	// 初始化第一个状态
	states[0][0] = 0
	if items[0] <= w {
		states[0][items[0]] = values[0]
	}

	for i := 1; i < 5; i++ {
		for j := 0; j < 10; j++ {
			if states[i-1][j] >= 0 {
				// 不放[i]物品
				if states[i-1][j] > states[i][j] {
					states[i][j] = states[i-1][j]
				}

				// 放物品[i], 但之前重量+当前重量不超限，同时在多种可能中，保留价值最大的
				if j+items[i] <= w {
					if values[i]+states[i-1][j] > states[i][j+items[i]] {
						states[i][j+items[i]] = values[i] + states[i-1][j]
					}
				}
			}
		}
	}

	for j := w; j >= 0; j-- {
		fmt.Println(j, states[4][j])
	}

	// 逆向判断所装物品
	j := 9
	for i := 4; i > 0; i-- {
		if j-items[i] >= 0 && states[i-1][j-items[i]] >= 0 {
			fmt.Println("buy: ", items[i])
			j -= items[i]
			fmt.Println("j===:", j)
		}
	}
	if j != 0 {
		fmt.Println("buy: ", items[0])
	}
}

func calValue2() {
	// 申请一维数组保存状态点
	items := [5]int{2, 2, 4, 6, 3}
	values := [5]int{3, 4, 8, 9, 6}
	w := 9

	var states [10]int
	for j := 0; j < 10; j++ {
		states[j] = -1
	}
	states[0] = 0
	if items[0] <= w {
		states[items[0]] = values[0]
	}
	fmt.Println(states)

	for i := 1; i < 5; i++ {
		for j := w - items[i]; j >= 0; j-- {
			if states[j] >= 0 {
				if states[j]+values[i] > states[j+items[i]] {
					states[j+items[i]] = states[j] + values[i]
				}
			}
		}
		fmt.Println(states)
	}

	for j := w; j >= 0; j-- {
		fmt.Println(j, states[j])
	}

}
