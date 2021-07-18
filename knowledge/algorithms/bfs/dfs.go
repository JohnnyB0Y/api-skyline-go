//  dfs.go
//
//
//  Created by JohnnyB0Y on 2021/07/18.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package bfs

import (
	"api-skyline-go/knowledge/custom_data_structures/stack"
	"fmt"
)

func DFS(graph map[string][]string, startPoint string) {
	// 访问栈
	s := stack.NewStack(40)
	s.Push(startPoint)

	// 存放已访问的节点
	seen := make(map[string]bool, len(graph))
	seen[startPoint] = true

	for !s.Empty() {
		vertex, err := s.Pop()
		if err != nil {
			fmt.Println("err:", err)
		}

		nodes := graph[vertex]
		for _, w := range nodes {
			// 未访问
			if see := seen[w]; !see {
				s.Push(w)
				seen[w] = true
			}
		}
		fmt.Println(vertex)
	}
}
