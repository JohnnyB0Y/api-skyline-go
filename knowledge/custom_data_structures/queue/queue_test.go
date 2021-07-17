//  queue_test.go
//
//
//  Created by JohnnyB0Y on 2021/07/17.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue([]int{3, 4, 2, 1})

	fmt.Println(q.Dequeue())

	for i := 0; i < 2; i++ {
		err := q.Enqueue(i + 20)
		result := "✔️"
		if err != nil {
			result = "×"
		}
		fmt.Printf("enqueue %d, %v\n", i+20, result)
	}

	for {
		val, err := q.Dequeue()
		if err != nil {
			break
		}
		fmt.Println("dequeue:", val)
	}

	i := 0
	for {
		err := q.Enqueue(i)
		if err != nil {
			break
		}
		fmt.Println("enqueue:", i)
		i += 6
	}

	// [ 0 6 12 18 ] 4
	queuePrintAll(&q)

}

func TestPriorityQueue(t *testing.T) {
	pq := NewPriorityQueue([]int{3, 4, 2, 1})
	fmt.Println(pq.Dequeue())

	for i := 0; i < 2; i++ {
		err := pq.Enqueue(i + 20)
		result := "✔️"
		if err != nil {
			result = "×"
		}
		fmt.Printf("enqueue %d, %v\n", i+20, result)
	}

	for {
		val, err := pq.Dequeue()
		if err != nil {
			break
		}
		fmt.Println("dequeue:", val)
	}

	i := 0
	for {
		err := pq.Enqueue(i)
		if err != nil {
			break
		}
		fmt.Println("enqueue:", i)
		i += 6
	}

	// [ 0 6 12 18 ] 4
	priorityQueuePrintAll(&pq)
}

func queuePrintAll(q *queue) {
	q.Enumerate(func(idx, val int) (stop bool) {
		if idx == 0 {
			fmt.Print("[ ")
		}
		fmt.Printf("%d ", val)
		if idx == q.length-1 {
			fmt.Println("]", q.length)
		}
		return
	})
}

func priorityQueuePrintAll(pq *PriorityQueue) {
	pq.Enumerate(func(idx, val int) (stop bool) {
		if idx == 0 {
			fmt.Print("[ ")
		}
		fmt.Printf("%d ", val)
		return
	})
	fmt.Println("]")
}
