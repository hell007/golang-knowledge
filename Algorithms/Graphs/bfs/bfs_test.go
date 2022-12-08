/*
 * @Descripttion:
 * @Author: zenghua.wang
 * @Date: 2021-02-21 21:29:33
 * @LastEditors: zenghua.wang
 * @LastEditTime: 2022-12-08 14:30:40
 */
package bfs

import (
	"golang-knowledge/Data-Structures/graph"
	"testing"

	"github.com/stretchr/testify/assert"
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
