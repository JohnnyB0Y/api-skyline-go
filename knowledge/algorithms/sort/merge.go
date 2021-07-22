//  merge.go
//
//
//  Created by JohnnyB0Y on 2021/07/21.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package sort

import (
	"sync"
)

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
func MergeSortWithMultiGoroutine(nums []int, goroutines int) []int {

	if len(nums) < 3 {
		return MergeSort(nums)
	}

	var wg sync.WaitGroup
	wg.Add(goroutines)

	// 按goroutine数切分
	fragment := len(nums) / goroutines
	results := make([][]int, goroutines)
	for g := 0; g < goroutines; g++ {
		go func(g int) {
			start := g * fragment
			end := start + fragment
			if g == goroutines-1 { // 最后一个
				end = len(nums)
			}
			results[g] = MergeSort(nums[start:end]) // 记录每条 goroutine完成的结果
			wg.Done()
		}(g)
	}

	wg.Wait()

	// 合并, 4 x 4 <=等价=> 2 x 2 x 2 x 2
	nums = results[0] // 取出第一个
	for i := 1; i < len(results); i++ {
		nums = merge(nums, results[i])
	}

	return nums
}
