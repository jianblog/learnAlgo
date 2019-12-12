package sort

/*
    Author: Liujianbo
	Date: 20191207
	Desc: 二分查找法及若干变种问题
*/
import (
	_ "fmt"
)

func bSearch(arr []int, v int) int {
	// 二分查找
	n := len(arr)
	lo, hi := 0, n-1
	for lo <= hi {
		mid := lo + ((hi - lo) >> 1)
		if arr[mid] == v {
			return mid
		}
		if arr[mid] > v {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return -1
}

func binaryOtherSearch1(arr []int, v int) int {
	// 二分查找 变种问题1: 查找第一个等于给定值元素
	n := len(arr)
	lo, hi := 0, n-1
	for lo <= hi {
		mid := lo + ((hi - lo) >> 1)
		if arr[mid] > v {
			hi = mid - 1
		} else if arr[mid] < v {
			lo = mid + 1
		} else {
			for mid >= 1 && arr[mid-1] == v {
				mid -= 1
			}
			return mid
			/* 优化样例: 替换else 内部
			if mid == 0 || arr[mid-1] != v {
				return mid
			} else {
				high = mid - 1
			}
			*/
		}
	}
	return -1
}

func binaryOtherSearch2(arr []int, v int) int {
	// 二分查找 变种问题2：查找最后一个等于给定值得元素
	n := len(arr)
	lo, hi := 0, n-1
	for lo <= hi {
		mid := lo + ((hi - lo) >> 1)
		if arr[mid] > v {
			hi = mid - 1
		} else if arr[mid] < v {
			lo = mid + 1
		} else {
			for mid <= n-2 && arr[mid+1] == v {
				mid += 1
			}
			return mid
			/* 优化样例：替换else内部
			if mid == n -1 || arr[mid+1] != v {
				return mid
			} else {
				lo = mid + 1
			}
			*/
		}
	}
	return -1
}

func binaryOtherSearch3(arr []int, v int) int {
	// 二分查找 变种问题3：查找第一个大于等于给定值得元素
	n := len(arr)
	lo, hi := 0, n-1
	for lo <= hi {
		mid := lo + ((hi - lo) >> 1)
		if arr[mid] >= v {
			if mid == 0 || arr[mid-1] < v {
				return mid
			} else {
				hi = mid - 1
			}
		} else {
			lo = mid + 1
		}
	}
	return -1
}

func binaryOtherSearch4(arr []int, v int) int {
	// 二分查找 变种问题4：查找最后一个小于等于给定值的元素
	n := len(arr)
	lo, hi := 0, n-1
	for lo <= hi {
		mid := lo + ((hi - lo) >> 1)
		if arr[mid] <= v {
			if mid == n-1 || arr[mid+1] > v {
				return mid
			} else {
				lo = mid + 1
			}
		} else {
			hi = mid - 1
		}
	}
	return -1
}
