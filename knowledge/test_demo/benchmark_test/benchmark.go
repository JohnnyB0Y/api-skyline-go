//  benchmark.go
//
//
//  Created by JohnnyB0Y on 2021/07/19.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package benchmarktest

func MakeSliceWithoutPreallocatedMemory(size int) []int {
	var s []int

	for i := 0; i < size; i++ {
		s = append(s, i)
	}
	return s
}

func MakeSliceWithPreallocatedMemory(size int) []int {
	var s = make([]int, 0, size)

	for i := 0; i < size; i++ {
		s = append(s, i)
	}
	return s
}
