/*
 * @Descripttion:
 * @Author: zenghua.wang
 * @Date: 2021-02-21 21:29:33
 * @LastEditors: zenghua.wang
 * @LastEditTime: 2022-12-08 14:38:12
 */
package dijkstra

import (
	"fmt"
	"golang-knowledge/Data-Structures/graph"
	"testing"
)

func TestShortestPath(t *testing.T) {
	h := graph.NewUndirected()

	for i := 0; i < 5; i++ {
		h.AddVertex(graph.VertexId(i))
	}

	h.AddEdge(graph.VertexId(0), graph.VertexId(1), 10)
	h.AddEdge(graph.VertexId(1), graph.VertexId(2), 20)
	h.AddEdge(graph.VertexId(2), graph.VertexId(3), 40)
	h.AddEdge(graph.VertexId(0), graph.VertexId(2), 50)
	h.AddEdge(graph.VertexId(0), graph.VertexId(3), 80)
	h.AddEdge(graph.VertexId(0), graph.VertexId(4), 10)
	h.AddEdge(graph.VertexId(4), graph.VertexId(3), 10)

	prev := ShortestPath(h, graph.VertexId(0))
	fmt.Println(prev)
	if prev[1] != graph.VertexId(0) ||
		prev[2] != graph.VertexId(1) ||
		prev[3] != graph.VertexId(4) ||
		prev[4] != graph.VertexId(0) {

		fmt.Println(prev)
		t.Error()
	}
}
