//  queue.go
//
//
//  Created by JohnnyB0Y on 2021/07/17.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package queue

import (
	"errors"
)

type queue struct {
	head   int
	tail   int
	arr    []int
	length int
}

func NewQueue(arr []int) queue {
	q := queue{}
	q.length = len(arr)
	q.head = 0
	q.tail = 0 // 此时队列满了
	q.arr = arr
	return q
}

func (q *queue) Empty() bool {
	return q.head == -1
}

func (q *queue) Full() bool {
	return q.tail == q.head
}

func (q *queue) Enqueue(item int) error {
	// 已满
	if q.Full() {
		return errors.New("is full baby")
	}

	// 首次进入
	if q.head == -1 {
		q.head = 0
	}

	// 入队
	q.arr[q.tail] = item
	q.tail = q.nextTail()

	q.length += 1
	return nil
}

func (q *queue) Dequeue() (int, error) {
	if q.Empty() {
		return 0, errors.New("is empty baby")
	}

	item := q.arr[q.head]
	q.arr[q.head] = 0
	q.head = q.nextHead()

	// 相等，就回到过去
	if q.head == q.tail {
		q.head = -1
		q.tail = 0
	}

	q.length -= 1
	return item, nil
}

// 枚举遍历堆
func (q *queue) Enumerate(fn func(idx, val int) (stop bool)) {
	if q.Empty() {
		return
	}
	if q.tail > q.head {
		for i := q.head; i < q.tail; i++ {
			stop := fn(i, q.arr[i])
			if stop {
				break
			}
		}
	} else {
		for i := q.head; i < len(q.arr); i++ {
			stop := fn(i, q.arr[i])
			if stop {
				break
			}
		}
		for i := 0; i < q.tail; i++ {
			stop := fn(i, q.arr[i])
			if stop {
				break
			}
		}
	}
}

func (q *queue) nextTail() int {
	if q.tail+1 == len(q.arr) {
		return 0
	}
	return q.tail + 1
}

func (q *queue) nextHead() int {
	if q.head+1 == len(q.arr) {
		return 0
	}
	return q.head + 1
}
