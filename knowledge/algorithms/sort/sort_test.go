//  sort_test.go
//
//
//  Created by JohnnyB0Y on 2021/07/20.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	nums := []int{14, 13, 11, 18, 19, 24, 3, 5, 1}
	fmt.Println("start:", nums)
	QuickSort(nums)
	// [1 3 5 11 13 14 18 19 24]
	fmt.Println("end", nums)

	nums = []int{1, 2, 89, 43, 23, 11, 0, 34, 99, 100, 99, 2}
	fmt.Println("start:", nums)
	QuickSort(nums)
	// [0 1 2 2 11 23 34 43 89 99 99 100]
	fmt.Println("end", nums)

	nums = []int{1, 2, -89, 43, 23, -11, 0, 34, 99, 100, -99, 2}
	fmt.Println("start:", nums)
	QuickSort(nums)
	// [-99 -89 -11 0 1 2 2 23 34 43 99 100]
	fmt.Println("end", nums)
}

func TestMergeSort(t *testing.T) {
	nums := []int{14, 13, 11, 18, 19, 24, 3, 5, 1}
	fmt.Println("start:", nums)
	nums = MergeSort(nums)
	// [1 3 5 11 13 14 18 19 24]
	fmt.Println("end", nums)

	nums = []int{1, 2, 89, 43, 23, 11, 0, 34, 99, 100, 99, 2}
	fmt.Println("start:", nums)
	nums = MergeSort(nums)
	// [0 1 2 2 11 23 34 43 89 99 99 100]
	fmt.Println("end", nums)

	nums = []int{1, 2, -89, 43, 23, -11, 0, 34, 99, 100, -99, 2}
	fmt.Println("start:", nums)
	nums = MergeSort(nums)
	// [-99 -89 -11 0 1 2 2 23 34 43 99 100]
	fmt.Println("end", nums)
}
