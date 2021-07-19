//  heap_test.go
//
//
//  Created by JohnnyB0Y on 2021/07/17.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package heap

import (
	"api-skyline-go/knowledge/custom_data_structures/item"
	"fmt"
	"testing"
)

type MyInt struct {
	val int
}

func (i MyInt) ComparisonKey() int {
	return i.val
}

func TestHeap(t *testing.T) {
	ints := []item.ItemContainer{
		MyInt{1}, MyInt{3}, MyInt{5}, MyInt{7}, MyInt{15}, MyInt{0}, MyInt{17},
	}

	h := NewHeap(ints)
	heapPrintAll(&h)
	heapPopAll(&h)

	ints2 := []item.ItemContainer{
		MyInt{1}, MyInt{2}, MyInt{10}, MyInt{3}, MyInt{17},
	}
	h.Build(ints2)
	heapPrintAll(&h)
	v1, _ := h.Pop()
	v2, _ := h.Pop()
	fmt.Printf("v1 %v v2 %v\n", v1, v2)
	heapPrintAll(&h)

	h.Push(MyInt{50})
	h.Push(MyInt{11})
	heapPrintAll(&h)
	heapPopAll(&h)
}

func heapPopAll(h *Heap) {
	fmt.Print("[")
	for {
		val, err := h.Pop()
		if err != nil {
			break
		}
		fmt.Printf("%d ", val)
	}
	fmt.Println("]")
}

func heapPrintAll(h *Heap) {
	h.Enumerate(func(idx int, val item.ItemContainer) (stop bool) {
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
