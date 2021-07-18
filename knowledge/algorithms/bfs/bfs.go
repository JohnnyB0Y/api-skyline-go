//  bfs.go
//
//
//  Created by JohnnyB0Y on 2021/07/18.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package bfs

import (
	"api-skyline-go/knowledge/custom_data_structures/queue"
	"fmt"
)

func BFS(graph map[string][]string, startPoint string) map[string]string {
	// 访问队列
	q := queue.NewQueue(make([]string, 0, len(graph)), len(graph))
	q.Enqueue(startPoint)

	// 存放已访问的节点
	seen := make(map[string]bool, len(graph))
	seen[startPoint] = true
	// 存放最短路径
	pathMap := make(map[string]string, len(graph))
	pathMap[startPoint] = ""

	for !q.Empty() {
		vertex, err := q.Dequeue()
		if err != nil {
			fmt.Println("err:", err)
			break
		}

		// nodes := graph[vertex.(string)]
		nodes := graph[vertex]
		for _, w := range nodes {
			// 未访问
			if see := seen[w]; !see {
				q.Enqueue(w)
				seen[w] = true
				pathMap[w] = vertex
			}
		}
	}
	return pathMap
}

// 两点间最短路径
func ShortestPathP2P(graph map[string][]string, start, end string) (paths []string) {
	pathMap := BFS(graph, start)
	for end != "" {
		paths = append(paths, end)
		end = pathMap[end]
	}
	for i, j := 0, len(paths)-1; i < j; i, j = i+1, j-1 {
		paths[i], paths[j] = paths[j], paths[i]
	}
	return
}
