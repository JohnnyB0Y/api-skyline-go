//  quick.go
//
//
//  Created by JohnnyB0Y on 2021/07/20.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package sort

import "fmt"

func QuickSort(nums []int) {
	var q [][]int = make([][]int, 0, 10)
	q = append(q, []int{0, len(nums) - 1})

	for len(q) > 0 {
		fragment := q[0]
		q = q[1:] // 模拟队列弹出操作
		low := fragment[0]
		hight := fragment[1]

		fmt.Println("start:", fragment, q)
		if hight > low { // hight=low 表示一个元素，hight>low 表示两个及以上，一个元素跳过不处理
			N := hight - low + 1 // hight=1 low=0 两个元素；hight>1 low=0 三个及以上;
			if N == 2 {          // 刚好两个数，处理并跳过！！！
				if nums[low] > nums[hight] { // 如果 low > hight 直接交换
					swap(&nums, low, hight)
				}
				continue
			}
			// 切割
			pivot := partition(&nums, low, hight)
			q = append(q, []int{low, pivot - 1})
			q = append(q, []int{pivot + 1, hight})
		}
	}
}

func partition(nums *[]int, low, hight int) int {
	pivot := low
	start := low + 1

	// 找中心轴，不要让pivot是最小值
	if (*nums)[pivot] < (*nums)[start] {
		swap(nums, pivot, start)
	}

	for i := start; i <= hight; i++ {
		if (*nums)[i] < (*nums)[pivot] {
			swap(nums, i, start)
			start += 1
		}
	}

	swap(nums, pivot, start-1)
	return start - 1
}

func swap(nums *[]int, low, hight int) {
	temp := (*nums)[low]
	(*nums)[low] = (*nums)[hight]
	(*nums)[hight] = temp
}
