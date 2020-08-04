package topological

import (
	"fmt"
	"github.com/hell007/golang-knowledge/Data-Structures/graph"
	"testing"
)

func TestSort(t *testing.T) {
	h := graph.NewDirected()

	h.AddVertex(graph.VertexId(2))
	h.AddVertex(graph.VertexId(3))
	h.AddVertex(graph.VertexId(5))
	h.AddVertex(graph.VertexId(7))
	h.AddVertex(graph.VertexId(8))
	h.AddVertex(graph.VertexId(9))
	h.AddVertex(graph.VertexId(10))
	h.AddVertex(graph.VertexId(11))

	h.AddEdge(graph.VertexId(7), graph.VertexId(8), 1)
	h.AddEdge(graph.VertexId(7), graph.VertexId(11), 1)
	h.AddEdge(graph.VertexId(5), graph.VertexId(11), 1)
	h.AddEdge(graph.VertexId(3), graph.VertexId(8), 1)
	h.AddEdge(graph.VertexId(3), graph.VertexId(10), 1)
	h.AddEdge(graph.VertexId(8), graph.VertexId(9), 1)
	h.AddEdge(graph.VertexId(11), graph.VertexId(10), 1)
	h.AddEdge(graph.VertexId(11), graph.VertexId(2), 1)
	h.AddEdge(graph.VertexId(11), graph.VertexId(9), 1)

	s := Sort(h)
	i, _ := s.Pop()
	first := graph.VertexId(i)
	if first != 7 &&
		first != 5 &&
		first != 3 {
		fmt.Print(first, first != 7)
		t.Error()
	}

}
