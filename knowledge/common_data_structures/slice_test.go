//  slice_test.go
//
//
//  Created by JohnnyB0Y on 2021/07/04.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package commondatastructures_test

import (
	"fmt"
	"testing"
)

func sliceRise(s []int) {
	s = append(s, 0)
	for i := range s {
		// s[i]++  <=============>  s[i] = s[i] + s[i]
		s[i]++
	}
}

func slicePrint() {
	s1 := []int{1, 2}
	s2 := s1
	s2 = append(s2, 3)

	sliceRise(s1)
	sliceRise(s2)
	fmt.Println(s1, s2)
}

func sliceExtend() {
	var slice []int

	s1 := append(slice, 1, 2, 3)
	s11 := append(slice, 1, 2, 3, 4)
	s12 := append(slice, 1, 2, 3, 4, 5)
	// 0 4 4 6
	fmt.Println(cap(slice), cap(s1), cap(s11), cap(s12))
	s2 := append(s1, 4)
	fmt.Println(&s1[0] == &s2[0])

	s13 := append(s1, 4, 5)
	// 4 8
	fmt.Println(cap(s1), cap(s13))

	s14 := make([]int, 999)
	s15 := append(s14, 0)
	s16 := append(s14, 0, 1, 2)
	s17 := append(s14, s14...) // len 2000
	s17 = append(s17, s14...)  // len 3000
	// 999 2048 2048 3408
	fmt.Println(cap(s14), cap(s15), cap(s16), cap(s17))

	// 当切片从0容量开始扩容时，会以偶数形式扩容。例如 长度需要3，那么容量会分配4。长度需要5，容量会分配6。
	// 当切片容量大于0时进行扩容，会以当前容量的2倍进行扩容。例如 当前容量为4，那么容量会扩容到8。
	// 当切片容量较大时，扩容的规律不可靠，所以不要依赖实现，要依赖抽象。

	//
	// 指定切片的容量，减少扩容时的内存分配和拷贝次数；
	//
}

func sliceFunny() {
	s := *new([]int)
	fmt.Println(s, s == nil)
}

func TestSlice(t *testing.T) {
	slicePrint()
	sliceExtend()
	sliceFunny()
}
