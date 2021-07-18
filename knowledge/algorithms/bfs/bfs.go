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

func BFS(graph map[string][]string, startPoint string) {
	// 访问队列
	q := queue.NewQueue(make([]string, 0, len(graph)), 40)
	q.Enqueue(startPoint)

	// 存放已访问的节点
	seen := make(map[string]bool, len(graph))
	seen[startPoint] = true

	for !q.Empty() {
		vertex, err := q.Dequeue()
		if err != nil {
			fmt.Println("err:", err)
		}

		// nodes := graph[vertex.(string)]
		nodes := graph[vertex]
		for _, w := range nodes {
			// 未访问
			if see := seen[w]; !see {
				q.Enqueue(w)
				seen[w] = true
			}
		}
		fmt.Println(vertex)
	}
}
