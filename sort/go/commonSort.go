package sort

func BubbleSort(arr []int, n int) []int {
	// 冒泡排序
	if n == 0 || n == 1 {
		return arr
	}
	for i := 0; i < n; i++ {
		flag := 0
		for j := 1; j < n-i; j++ {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
				flag = 1
			}
		}
		// 有序状态时，不会有任何移动操作
		if flag == 0 {
			return arr
		}
	}
	return arr
}

func InsertionSort(arr []int, n int) []int {
	// 插入排序
	if n == 0 || n == 1 {
		return arr
	}
	for i := 1; i < n; i++ {
		p := arr[i]
		j := i - 1
		for ; j >= 0; j-- {
			if p < arr[j] {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		// j循环先自减然后比较
		arr[j+1] = p
	}
	return arr
}

func SelectionSort(arr []int, n int) []int {
	// 选择排序
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}

func MergeSort(arr []int, l int, r int) []int {
	//归并排序： 拆分-->合并
	if l >= r {
		return arr[l : r+1]
	} else {
		m := l + (r-l)/2
		ret := pJoin(MergeSort(arr, l, m), MergeSort(arr, m+1, r))
		return ret
	}
}
func pJoin(arrA []int, arrB []int) []int {
	// 合并两个切片
	var ret []int
	p1, p2 := 0, 0
	lenA, lenB := len(arrA), len(arrB)
	for p1 < lenA && p2 < lenB {
		if arrA[p1] < arrB[p2] {
			ret = append(ret, arrA[p1])
			p1 += 1
		} else {
			ret = append(ret, arrB[p2])
			p2 += 1
		}
	}
	if p1 == lenA {
		ret = append(ret, arrB[p2:]...)
	}
	if p2 == lenB {
		ret = append(ret, arrA[p1:]...)
	}
	return ret
}
