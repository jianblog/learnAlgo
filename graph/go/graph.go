package graph

import (
	"container/list"
	"fmt"
)

type Graph struct {
	adj []*list.List
	v   int
}

func newGraph(v int) *Graph {
	graph := &Graph{}
	graph.v = v
	graph.adj = make([]*list.List, 3)
	for i := range graph.adj {
		graph.adj[i] = list.New()
	}
	return graph
}
