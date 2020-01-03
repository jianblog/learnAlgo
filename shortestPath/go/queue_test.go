package shortestPath

import (
	"fmt"
	"testing"
)

func printPath(s int, t int, path []int) {
	if s == t {
		return
	}
	printPath(s, path[t], path)
	fmt.Printf("->%v", t)
}

func TestDijkstra(t *testing.T) {
	myGraph := newGraph(6)
	myGraph.addEdge(0, 1, 10)
	myGraph.addEdge(0, 4, 15)
	myGraph.addEdge(1, 2, 15)
	myGraph.addEdge(1, 3, 2)
	myGraph.addEdge(2, 5, 5)
	myGraph.addEdge(3, 2, 1)
	myGraph.addEdge(3, 5, 12)
	myGraph.addEdge(4, 5, 10)

	path := dijkstra(*myGraph, 0, 5)
	fmt.Printf("%v", 0)
	printPath(0, 5, path)
	fmt.Printf("\n")

	/*
		e := myGraph.adj[3]
		for i := e.Front(); i != nil; {
			fmt.Printf("%v\n", i.Value.(Edge).sid)
			i = i.Next()
		}
	*/
}
