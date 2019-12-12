package heap

/*
堆：使用数组模拟堆(完全二叉树)
第i位的左右子节点 (i+1)*2 -1, (i+1)*2
*/
import (
	_ "fmt"
)

type Heap struct {
	// 为便于堆定位计算，从[1]开始保存数据，[0]空闲不用
	arr   []int
	size  int
	count int
}

func NewHeap(capacity int) *Heap {
	heap := &Heap{make([]int, capacity+1), capacity, 0}
	return heap
}

func (this *Heap) Insert(v int) bool {
	if this.count == 0 {
		this.arr[1] = v
		this.count += 1
		return true
	} else if this.count == this.size {
		return false
	} else {
		// 从数组最末尾开始向上比较
		pos := this.count + 1
		this.arr[pos] = v
		for i := pos / 2; i > 0; {
			if this.arr[i] >= this.arr[pos] {
				this.count += 1
				return true
			} else {
				this.arr[i], this.arr[pos] = this.arr[pos], this.arr[i]
				pos = i
				i = i / 2
			}
		}
		return true
	}
}

func (this *Heap) DeleteTop() bool {
	// 删除顶部节点, 最末元素置顶，再比较交换
	if this.count == 0 {
		return false
	} else if this.count == 1 {
		this.count = 0
		return true
	} else {
		this.arr[1] = this.arr[this.count]
		this.count -= 1
		maxpos := 1
		for i := 1; i < this.count; {
			if this.arr[i] < this.arr[i*2] {
				maxpos = i * 2
			}
			if i*2+1 <= this.count && this.arr[i*2+1] > this.arr[maxpos] {
				maxpos = i*2 + 1
			}
			if maxpos != i {
				this.arr[i], this.arr[maxpos] = this.arr[maxpos], this.arr[i]
				i = maxpos
			} else {
				break
			}
		}
		return true
	}
}
