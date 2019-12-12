package heap

/*
堆：使用数组模拟堆(完全二叉树)
数组索引位对应关系：
    1. 第i位的左右子节点 i*2 + 1, (i+1)*2
	2. 第i位的父节点为   (i-1)/2   i>0
*/
import (
	_ "fmt"
)

type Heap struct {
	arr   []int
	size  int
	count int
}

func NewHeap(capacity int) *Heap {
	heap := &Heap{make([]int, capacity), capacity, 0}
	return heap
}

func (this *Heap) Insert(v int) bool {
	if this.count == 0 {
		this.arr[0] = v
		this.count += 1
		return true
	} else if this.count == this.size {
		return false
	} else {
		// 从数组最末尾开始向上比较
		pos := this.count
		this.count += 1
		this.arr[pos] = v
		for i := (pos - 1) / 2; i >= 0; {
			if this.arr[i] >= this.arr[pos] {
				return true
			} else {
				this.arr[i], this.arr[pos] = this.arr[pos], this.arr[i]
				pos = i
				i = (i - 1) / 2
			}
		}
		return true
	}
}

func (this *Heap) DeleteTop() bool {
	// 删除顶部节点, 最末元素置顶，再比较交换
	if this.count == 0 {
		return false
	} else {
		this.arr[0] = this.arr[this.count-1]
		this.count -= 1
		HeapifyToTail(this.arr, this.count)
		return true
	}
}
func HeapifyToTail(arr []int, count int) {
	maxpos := 0
	for i := 0; i < count; {
		if i*2+1 < count && arr[i] < arr[i*2+1] {
			maxpos = i*2 + 1
		}
		if (i+1)*2 < count && arr[(i+1)*2] > arr[maxpos] {
			maxpos = (i + 1) * 2
		}
		if maxpos != i {
			arr[i], arr[maxpos] = arr[maxpos], arr[i]
			i = maxpos
		} else {
			break
		}
	}
}

func Heapify(arr []int) {
	//从底部对每个非叶节点开始进行调整
	// 倒数第一个非叶节点是 (i-1)/2
	length := len(arr)
	for i := (length - 2) / 2; i >= 0; i-- {
		maxpos := i
		j := i
		for j < length {
			if j*2+1 < length && arr[j] < arr[j*2+1] {
				maxpos = j*2 + 1
			}
			if (j+1)*2 < length && arr[(j+1)*2] > arr[maxpos] {
				maxpos = (j + 1) * 2
			}
			if maxpos != j {
				arr[j], arr[maxpos] = arr[maxpos], arr[j]
				j = maxpos
			} else {
				break
			}
		}
	}
}
