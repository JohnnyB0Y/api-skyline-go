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

/**
旋转数组
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。



进阶：

尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2skh7/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func rotate(nums []int, k int) {
	// k = 3
	// [1,2,3,(4),5,6,7]
	// 1 -> 4; 4 -> 7; 7 -> 3; 3 -> 6; 6 -> 2; 2 -> 5; 5 -> 1;
	//
	// [5,6,7,(1),2,3,4]

	// k = 4
	// [1,2,3,4,(5),6,7]
	// 1 -> 5; 5 -> 2; 2 -> 6; 6 -> 3; 3 -> 7; 7 -> 4; 4 -> 1;
	// [4,5,6,7,(1),2,3]

	// k = 4
	// [1,2,3,4,(5),6,7,8]
	// 1 -> 5; 5 -> 1;
	// 2 -> 6; 6 -> 2;
	// 3 -> 7; 7 -> 3;
	// 4 -> 8; 8 -> 4;
	// [5,6,7,8,(1),2,3,4]

	// k = 3
	// [1,2,3,(4),5,6,7,8,9]
	// 1 -> 4; 4 -> 7; 7 -> 1;
	// 2 -> 5; 5 -> 8; 8 -> 2;
	// 3 -> 6; 6 -> 9; 9 -> 3;
	// [7,8,9,(1),2,3,4,5,6]

	// k = 4
	// [1]
	// [1,2]

	// k = 4 (k > len(arr)) ==> k = k % len(arr) ==> k = 1
	// [1,2,3]
	// 1 -> 2; 2 -> 3; 3 ->1;
	// [3,1,2]

	// k = 4
	// [1,2,3,4,5,6]
	// 1 -> 5; 5 -> 3; 3 -> 1;
	// 2 -> 6; 6 -> 4; 4 -> 2;
	// [3,4,5,6,1,2]
	// [3,4,5,6,1,2]

	// 倒着走
	length := len(nums)

	if length < k {
		k = k % length
	}

	if k == 0 || length < 1 || length == k {
		return
	}

	temp := 0
	j := 0
	count := 0
	first := 0
	if length%k == 0 {
		// 整除
		for i := 0; i < k; i++ {
			j = length - k + i
			temp = nums[i]
			// nums[0] = nums[6]
			// nums[1] = nums[7]
			// nums[2] = nums[8]
			nums[i] = nums[j]
			for j > k+i {
				// nums[6] = nums[3]
				// nums[7] = nums[4]
				// nums[8] = nums[5]
				nums[j] = nums[j-k]
				j -= k
			}
			// nums[3] = temp
			// nums[4] = temp
			// nums[5] = temp
			nums[j] = temp
		}
	} else {
		// 不整除
		// k = 4
		// [1,2,3,4,(5),6,7]
		// 1 -> 5; 5 -> 2; 2 -> 6; 6 -> 3; 3 -> 7; 7 -> 4; 4 -> 1;
		// [4,5,6,7,(1),2,3]
		for count < length {
			j = length - k + first
			temp = nums[first]
			// nums[0] = nums[3]
			nums[first] = nums[j]
			count += 1
			for j != first {
				next := j - k
				if next < 0 {
					next = length + next
				}
				// nums[3] = nums[6]
				// nums[6] = nums[2]
				// nums[2] = nums[5]
				// nums[5] = nums[1]
				// nums[1] = nums[4]
				nums[j] = nums[next]
				count += 1
				if next == first { // 回到原点
					// nums[4] = nums[0]
					nums[j] = temp
					break
				}
				j = next
			}

			count += 1
			first += 1
		}
	}
}
