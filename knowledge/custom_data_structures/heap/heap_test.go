//  heap_test.go
//
//
//  Created by JohnnyB0Y on 2021/07/17.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package heap

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	h := NewHeap([]int{2, 3, 5, 7, 15, 0, 17})
	heapPrintAll(&h)
	heapPopAll(&h)

	h.Build([]int{1, 2, 10, 3, 17})
	heapPrintAll(&h)
	h.Pop()
	h.Pop()
	heapPrintAll(&h)

	h.Push(50)
	h.Push(11)
	heapPrintAll(&h)
	heapPopAll(&h)
}

func heapPopAll(h *heap) {
	for {
		val, err := h.Pop()
		if err != nil {
			break
		}
		fmt.Printf("%d ", val)
	}
	fmt.Println()
}

func heapPrintAll(h *heap) {
	h.Enumeration(func(idx, val int) (stop bool) {
		if idx == 0 {
			fmt.Print("[ ")
		}
		fmt.Printf("%d ", val)
		if idx == h.length-1 {
			fmt.Println("]", h.length)
		}
		return
	})
}
