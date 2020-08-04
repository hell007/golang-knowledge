package bfs

import (
	"github.com/hell007/golang-knowledge/Data-Structures/graph"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBfs(t *testing.T) {
	h := graph.NewDirected()

	for i := 0; i < 10; i++ {
		h.AddVertex(graph.VertexId(i))
	}

	for i := 0; i < 9; i++ {
		h.AddEdge(graph.VertexId(i), graph.VertexId(i+1), 1)
	}
	count := 0
	Bfs(h, graph.VertexId(3), func(id graph.VertexId) {
		count += int(id)
	})
	assert.Equal(t, 42, count)
}
