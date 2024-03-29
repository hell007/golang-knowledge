/*
 * @Descripttion:
 * @Author: zenghua.wang
 * @Date: 2021-02-21 21:29:33
 * @LastEditors: zenghua.wang
 * @LastEditTime: 2022-12-08 14:37:29
 */
package dfs

import (
	"golang-knowledge/Data-Structures/graph"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUndirectedDfs(t *testing.T) {
	//建立一个无向图
	h := graph.NewUndirected()

	//增加顶点，
	for i := 0; i < 10; i++ {
		h.AddVertex(graph.VertexId(i))
	}
	//增加边
	for i := 0; i < 9; i++ {
		h.AddEdge(graph.VertexId(i), graph.VertexId(i+1), 1)
	}
	counter := 0
	UndirectedDfs(h, graph.VertexId(4), func(id graph.VertexId) {
		counter += int(id)
	})
	assert.Equal(t, 45, counter)
}

func TestDirectedDfs(t *testing.T) {
	//建立一个有向图
	h := graph.NewDirected()

	for i := 0; i < 10; i++ {
		v := graph.VertexId(i)
		h.AddVertex(v)
	}

	for i := 0; i < 9; i++ {
		h.AddEdge(graph.VertexId(i), graph.VertexId(i+1), 1)
	}

	counter := 0
	DirectedDfs(h, graph.VertexId(3), func(v graph.VertexId) {
		counter += int(v)
	})

	assert.Equal(t, 42, counter)
}
