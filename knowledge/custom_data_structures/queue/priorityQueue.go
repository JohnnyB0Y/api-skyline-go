//  priorityQueue.go
//
//
//  Created by JohnnyB0Y on 2021/07/17.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package queue

import (
	"api-skyline-go/knowledge/custom_data_structures/heap"
	"api-skyline-go/knowledge/custom_data_structures/item"
)

// 优先队列
type PriorityQueue struct {
	h heap.Heap
}

func NewPriorityQueue(arr []item.ItemContainer) PriorityQueue {
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
func (pq *PriorityQueue) Enqueue(item item.ItemContainer) error {
	return pq.h.Push(item)
}

// 出队
func (pq *PriorityQueue) Dequeue() (item.ItemContainer, error) {
	return pq.h.Pop()
}

// 遍历
func (pq *PriorityQueue) Enumerate(fn func(idx int, val item.ItemContainer) (stop bool)) {
	pq.h.Enumerate(fn)
}
