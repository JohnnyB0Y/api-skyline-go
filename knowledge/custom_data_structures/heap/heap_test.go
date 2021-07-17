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

	fmt.Println(h.arr, h.length)

	for {
		val, err := h.Pop()
		if err != nil {
			break
		}
		fmt.Printf("%d ", val)
	}
	fmt.Println()
}
