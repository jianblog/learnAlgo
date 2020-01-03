package shortestPath

import (
	"fmt"
	"testing"
)

func TestCreateGraph(t *testing.T) {
	/* 初始化图测试
	   0->1 10, 0->4 15
	   1->2 15 1->3 2
	   2->5 5
	   3->2 1 3->5 12
	   4->5 10
	   5->
	*/
	myGraph := newGraph(6)
	myGraph.addEdge(0, 1, 10)
	myGraph.addEdge(0, 4, 15)
	myGraph.addEdge(1, 2, 15)
	myGraph.addEdge(1, 3, 2)
	myGraph.addEdge(2, 5, 5)
	myGraph.addEdge(3, 2, 1)
	myGraph.addEdge(3, 5, 12)
	myGraph.addEdge(4, 5, 10)

	fmt.Printf("%v\n", myGraph)
	fmt.Printf("%v\n", myGraph.adj[1].Front().Value)
}
