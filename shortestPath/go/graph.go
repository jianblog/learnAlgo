package shortestPath

/*
   图的数据结构，具有权值
   由外至内，分别涉及的数据结构为
   图 -> list.List -> Element -> Edge
*/

import (
	"container/list" // 双向链表
)

// 图的边，具有权值
type Edge struct {
	sid int
	tid int
	w   int
}

type Graph struct {
	adj []*list.List // 链表
	v   int
}

// 初始化指定顶点个数v的图，采用邻接表结构
func newGraph(v int) *Graph {
	graph := &Graph{}
	graph.v = v
	graph.adj = make([]*list.List, v)
	for i := range graph.adj {
		graph.adj[i] = list.New()
	}
	return graph
}

// 邻接表元素由带有权值的边组成
func (this *Graph) addEdge(s int, t int, w int) {
	this.adj[s].PushBack(Edge{s, t, w})
}
