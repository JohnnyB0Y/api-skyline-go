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

	fmt.Print(h.Pop())
	for 0 < h.length {
		fmt.Printf(" > %d", h.Pop())
	}
	fmt.Println()
}
