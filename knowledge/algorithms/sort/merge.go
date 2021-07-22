//  merge.go
//
//
//  Created by JohnnyB0Y on 2021/07/21.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package sort

// 归并排序
func MergeSort(nums []int) []int {
	N := len(nums)
	// 结束条件
	if N < 2 { // 0, 1
		return nums
	}

	// 分解
	left := MergeSort(nums[:N/2])
	right := MergeSort(nums[N/2:])
	// 合并
	return merge(left, right)
}

// 合并左右
func merge(left, right []int) []int {

	if len(left) < 1 || len(right) < 1 {
		return []int{}
	}
	if len(left) < 1 {
		return right
	}
	if len(right) < 1 {
		return left
	}

	nums := make([]int, 0, len(left)+len(right))
	l, r := 0, 0
	for i := 0; i < len(left)+len(right); i++ {

		if l >= len(left) || (r < len(right) && left[l] >= right[r]) { // 左边走完 或 左边比较大
			nums = append(nums, right[r])
			r++

		} else if r >= len(right) || (l < len(left) && left[l] <= right[r]) { // 右边走完 或 右边比较大
			nums = append(nums, left[l])
			l++
		}
	}
	return nums
}

// 多 goroutine 归并排序
func MergeSortByMultiGoroutine(nums []int) []int {

	return nums
}
