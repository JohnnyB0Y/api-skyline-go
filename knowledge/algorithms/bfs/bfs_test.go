//  bfs_test.go
//
//
//  Created by JohnnyB0Y on 2021/07/18.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package bfs

import "testing"

func TestBFS(t *testing.T) {
	graph := map[string][]string{
		"A": {"B", "C"},
		"B": {"A", "B", "C"},
		"C": {"A", "B", "D", "E"},
		"D": {"B", "C", "E", "F"},
		"E": {"C", "D"},
		"F": {"D"},
	}

	BFS(graph, "A")
}
