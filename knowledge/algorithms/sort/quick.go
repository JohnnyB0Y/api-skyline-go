//  quick.go
//
//
//  Created by JohnnyB0Y on 2021/07/20.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package sort

// 快速排序
func QuickSort(nums []int) {
	var q [][]int = make([][]int, 0, 10)
	q = append(q, []int{0, len(nums) - 1}) // 放入首个待处理的元素范围

	for len(q) > 0 {
		fragment := q[0] // 取出待处理的元素范围
		q = q[1:]        // 模拟队列弹出操作

		// fmt.Println("start:", "fragment:", fragment, "queue:", q)
		// 切割中心轴
		low, hight := fragment[0], fragment[1]
		pivot := partition(&nums, low, hight)

		// 处理左边
		fLLow, fLHight := low, pivot-1
		if fLHight-fLLow+1 == 2 && nums[fLLow] > nums[fLHight] { // 两个元素，如果 low > hight 直接交换
			swap(&nums, fLLow, fLHight)
		} else if fLHight > fLLow { // 三个元素以上，放入队列
			q = append(q, []int{fLLow, fLHight})
		}

		// 处理右边
		fRLow, fRHight := pivot+1, hight
		if fRHight-fRLow+1 == 2 && nums[fRLow] > nums[fRHight] { // 两个元素，如果 low > hight 直接交换
			swap(&nums, fRLow, fRHight)
		} else if fRHight > fRLow { // 三个元素以上，放入队列
			q = append(q, []int{fRLow, fRHight})
		}
	}
}

func partition(nums *[]int, low, hight int) int {
	// 找中心轴，不要让pivot是最小值
	pivot := low
	if (*nums)[pivot] < (*nums)[pivot+1] {
		swap(nums, pivot, pivot+1)
	}

	start := low + 2 // 上面已经处理了前两个元素，所以直接从第三个元素开始
	// 小于中心轴的，往左边交换
	for i := start; i <= hight; i++ {
		if (*nums)[i] < (*nums)[pivot] {
			swap(nums, i, start)
			start += 1 // 交换完，指向下一个待交换位置
		}
	}

	swap(nums, pivot, start-1) // start 指向下一个待交换位置，所以实际交换的位置需要 减一
	return start - 1
}

func swap(nums *[]int, low, hight int) {
	if low == hight {
		return
	}
	(*nums)[low], (*nums)[hight] = (*nums)[hight], (*nums)[low]
}

// 多 goroutine 快速排序
func QuickSortByMultiGoroutine(nums []int) {

}
