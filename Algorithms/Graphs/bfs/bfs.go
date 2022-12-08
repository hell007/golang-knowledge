/*
 * @Descripttion:
 * @Author: zenghua.wang
 * @Date: 2021-02-21 21:29:33
 * @LastEditors: zenghua.wang
 * @LastEditTime: 2022-12-08 14:10:34
 */
package bfs

import "golang-knowledge/Data-Structures/graph"

// 有向图的广度优先搜索
func Bfs(g *graph.DirGraph, start graph.VertexId, fn func(graph.VertexId)) {
	queue := []graph.VertexId{start}
	visited := make(map[graph.VertexId]bool)

	var next []graph.VertexId

	for len(queue) > 0 {
		next = []graph.VertexId{}
		for _, vertex := range queue {
			visited[vertex] = true
			//获取邻接点
			neighbours := g.GetNeighbours(vertex).VerticesIter()
			fn(vertex)
			for neighbour := range neighbours {
				if _, ok := visited[neighbour]; !ok {
					next = append(next, neighbour)
				}
			}
		}
		queue = next
	}
}
