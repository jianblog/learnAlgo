package shortestPath

/*
  Dijistra算法需要将顶点集合划分为两个不同集合queue,
  针对该集合内的元素数据结构为Vertex
*/
import (
	_ "fmt"
)

// queue中的元素类型
type Vertex struct {
	id   int
	dist int // 根据规则计算的最小距离
}

func NewVertex(id int, dist int) *Vertex {
	return &Vertex{id, dist}
}

type PriorityQueue struct {
	nodes []*Vertex // 保存为小顶堆结构
	count int       // 元素个数
}

func NewPriorityQueue(v int) *PriorityQueue {
	queue := &PriorityQueue{}
	queue.nodes = make([]*Vertex, 0)
	queue.count = 0
	return queue
}

// 小顶堆需要实现如下方法
func (this *PriorityQueue) heapify() {
	for i := (this.count - 2) / 2; i >= 0; i-- {
		minpos := i
		j := i
		// 每次循环都要遍历到末端
		for j < this.count {
			// i节点分别于左右子节点比较找出最小
			if j*2+1 < this.count && this.nodes[j].dist > this.nodes[j*2+1].dist {
				minpos = j*2 + 1
			}
			if (j+1)*2 < this.count && this.nodes[(j+1)*2].dist < this.nodes[minpos].dist {
				minpos = (j + 1) * 2
			}
			if minpos != j {
				this.nodes[i], this.nodes[minpos] = this.nodes[minpos], this.nodes[i]
				j = minpos
			} else {
				break
			}
		}
	}
}

// 取出最顶元素
func (this *PriorityQueue) poll() *Vertex {
	p := this.nodes[0]
	this.nodes[0] = this.nodes[this.count-1]
	this.count -= 1
	this.heapify()
	return p
}

func (this *PriorityQueue) add(vertex Vertex) {
	this.nodes = append(this.nodes, &vertex)
	this.count += 1
	this.heapify()
}

func (this *PriorityQueue) update(vertex Vertex) {
	for i := 0; i < this.count; i++ {
		if this.nodes[i].id == vertex.id {
			this.nodes[i].dist = vertex.dist
		}
	}
	this.heapify()
}

func (this *PriorityQueue) isEmpty() bool {
	if this.count == 0 {
		return true
	} else {
		return false
	}
}
