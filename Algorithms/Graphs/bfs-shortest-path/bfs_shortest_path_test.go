/*
 * @Descripttion:
 * @Author: zenghua.wang
 * @Date: 2021-02-21 21:29:33
 * @LastEditors: zenghua.wang
 * @LastEditTime: 2022-12-08 14:37:00
 */
package bfs_shortest_path

import (
	"golang-knowledge/Data-Structures/graph"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortestPath(t *testing.T) {
	h := graph.NewDirected()
	for i := 0; i < 10; i++ {
		h.AddVertex(graph.VertexId(i))
	}

	for i := 0; i < 9; i++ {
		h.AddEdge(graph.VertexId(i), graph.VertexId(i+1), 1)
	}

	assert.Equal(t, 9, GetDist(h, graph.VertexId(0), graph.VertexId(9)))
}
