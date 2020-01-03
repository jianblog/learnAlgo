package shortestPath

import (
	_ "fmt"
	"math"
)

func dijkstra(graph Graph, s int, t int) []int {
	// 初始化所有顶点的最短距离
	path := make([]int, graph.v)
	vertexes := make([]*Vertex, graph.v)
	for i := 0; i < graph.v; i++ {
		vertexes[i] = NewVertex(i, math.MaxInt32)
	}
	vertexes[s].dist = 0

	//待检测的顶点集合,小顶堆
	queue := NewPriorityQueue(graph.v)
	for i := 0; i < graph.v; i++ {
		queue.add(*vertexes[i])
	}

	for !queue.isEmpty() {
		minVertex := queue.poll()
		if minVertex.id == t {
			break
		}
		linkEdge := graph.adj[minVertex.id]
		for e := linkEdge.Front(); e != nil; {
			nextVertex := vertexes[e.Value.(Edge).tid]
			if minVertex.dist+e.Value.(Edge).w < nextVertex.dist {
				nextVertex.dist = minVertex.dist + e.Value.(Edge).w
				path[nextVertex.id] = minVertex.id
				queue.update(*nextVertex)
			}
			e = e.Next()
		}
	}
	return path
}
