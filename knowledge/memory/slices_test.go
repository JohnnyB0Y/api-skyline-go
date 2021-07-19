//  slices_test.go
//
//
//  Created by JohnnyB0Y on 2021/07/19.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package memory_test

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestEmptySlices(t *testing.T) {

	// nil 切片
	var s []int
	if s == nil {
		fmt.Printf("s 是 nil切片.\n")
	}

	// ---------------------------
	// 空切片
	p1 := make([]int, 0)
	p2 := make([]int, 0)

	// 对比空切片是否指向同一块内存地址？
	t.Log(p1, p2)
	t.Logf("p1 arr: %p, p2 arr: %p", p1, p2)

	// 类型和数据结构大小
	t.Logf("p1 Type: %T size: %v, p2 Type: %T size: %v", p1, unsafe.Sizeof(p1), p2, unsafe.Sizeof(p2))

	// 切片的长度和容量
	t.Logf("p1 len: %d cap: %d, p2 len: %d cap: %d", len(p1), cap(p1), len(p2), cap(p2))

	// ----------------------------
	// 指向非空数组的空切片
	arr := [...]int{1, 2, 3}

	s1 := arr[0:0:0]
	s2 := arr[0:0:0]

	// 对比空切片是否指向同一块内存地址？
	t.Log(s1, s2)
	t.Logf("\n s1 pointer: %p arr pointer: %p,\n s2 pointer: %p arr pointer: %p", &s1, s1, &s2, s2)
	t.Logf("arr pointer: %p", &arr)

	// 类型和数据结构大小
	t.Logf("s1 Type: %T size: %v, s2 Type: %T size: %v", s1, unsafe.Sizeof(s1), s2, unsafe.Sizeof(s2))

	// 切片的长度和容量
	t.Logf("s1 len: %d cap: %d, s2 len: %d cap: %d", len(s1), cap(s1), len(s2), cap(s2))

}

/**
声明但未初始化的切片 为 nil切片

初始化一个空切片，内部指向一个空数组
*/

// 切片 -> 栈
