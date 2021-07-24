//  array.go
//
//
//  Created by JohnnyB0Y on 2021/07/24.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

/**
删除排序数组中的重复项
给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2gy9m/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

package array

func RemoveDuplicates(nums []int) int {
	length := len(nums)
	if length < 2 {
		return length
	}

	cur, next := 0, 1
	for next < length {
		// 不相等
		if nums[cur] != nums[next] {
			for i := 1; i < length-next+1; i++ {
				nums[cur+i] = nums[next+i-1]
			}
			length -= next - cur - 1
			cur++
			next = cur + 1
			continue
		}
		// 相等，继续向前探
		next++

		// next走到最后 ? 把后面的重复项去除
		if next >= length {
			for i := 1; i < length-next+1; i++ {
				nums[cur+i] = nums[next+i-1]
			}
			length -= next - cur - 1
		}
	}
	return length
}

func RemoveDuplicates2(nums []int) int {
	length := len(nums)
	if length < 2 {
		return length
	}
	// fill 待填充的数组下标, find 寻找不同数字的数组下标
	fill, find := 1, 1
	for find < length {
		if nums[find-1] == nums[find] { // 相等，继续找
			find++
		} else { // 不相等
			nums[fill] = nums[find]
			fill++
			find++
		}
	}
	return fill
}

/**
买卖股票的最佳时机 II
给定一个数组 prices ，其中 prices[i] 是一支给定股票第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。



作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2zsx1/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func MaxProfit(prices []int) int {
	// 遇到低就买入，遇到最高就卖出
	// 寻找下一个低点，循环
	length := len(prices)
	if length < 2 {
		return 0
	}
	//  0 1 2 3 4 5
	// [7,1,5,3,6,4]
	buy, sale, profit := 0, 0, 0
	for buy < length-1 {
		// 当前价格高于下一个价格
		if prices[buy] > prices[buy+1] {
			buy++ // 继续找买入点
			continue
		}

		// 走到这里，算是找到了买入点，开始找卖出点
		sale = buy + 1
		for sale < length {
			if sale == length-1 || prices[sale] > prices[sale+1] { // sale 是最后一个 或 比下一个价格高 就不找了
				profit += prices[sale] - prices[buy]
				buy = sale + 1 // 下一轮
				break
			}
			sale++
		}
	}

	return profit
}
