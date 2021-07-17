//  priorityQueue.go
//
//
//  Created by JohnnyB0Y on 2021/07/17.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package queue

import "api-skyline-go/knowledge/custom_data_structures/heap"

// 优先队列
type PriorityQueue struct {
	h heap.Heap
}

func NewPriorityQueue(arr []int) PriorityQueue {
	pq := PriorityQueue{}
	pq.h = heap.NewHeap(arr)
	return pq
}

func (pq *PriorityQueue) Empty() bool {
	return pq.h.Empty()
}

func (pq *PriorityQueue) Full() bool {
	return pq.h.Full()
}

// 入队
func (pq *PriorityQueue) Enqueue(item int) error {
	return pq.h.Push(item)
}

// 出队
func (pq *PriorityQueue) Dequeue() (int, error) {
	return pq.h.Pop()
}

// 遍历
func (pq *PriorityQueue) Enumerate(fn func(idx, val int) (stop bool)) {
	pq.h.Enumerate(fn)
}
