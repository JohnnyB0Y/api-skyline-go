//  dp.go
//
//
//  Created by JohnnyB0Y on 2021/07/26.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package dp

/**
爬楼梯
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn854d/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func climbStairs(n int) int {

	return 0
}

/**
买卖股票的最佳时机
给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。

你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。

返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn8fsh/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func maxProfit(prices []int) int {
	if len(prices) < 1 {
		return 0
	}

	buy := prices[0]
	money := 0

	// 7, 1, 5, 3, 6, 4
	// 7: money = 0 buy = 7
	// 1: money = 0 buy = 1
	// 5: money = 4 buy = 1
	// 3: money = 4 buy = 1
	// 6: money = 5 buy = 1
	// 4: money = 5 buy = 1

	for i := 1; i < len(prices); i++ {
		// 看卖不卖
		if money < prices[i]-buy {
			money = prices[i] - buy
		}
		// 寻找最低点
		if prices[i] < buy {
			buy = prices[i]
		}
	}

	return money
}

/**
最大子序和
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
*/

func maxSubArray(nums []int) int {

	if len(nums) < 1 {
		return 0
	}
	max := nums[0] // 第一个必须选
	sum := max
	for i := 1; i < len(nums); i++ {
		// 更新 sum = Max{ nums[i], sum + nums[i] }
		if nums[i] > sum+nums[i] {
			sum = nums[i]
		} else {
			sum = sum + nums[i]
		}
		// 更新 max
		if max < sum {
			max = sum
		}
	}
	return max
}

/**
打家劫舍
你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnq4km/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func rob(nums []int) int {

	if len(nums) < 1 {
		return 0
	}

	// 间隔元素，取最大和
	total := []int{0, nums[0]} // 0，1

	for i := 1; i < len(nums); i++ {
		// total 填充了一个无选择时的状态 总价值为 0，所以这里的 i 和total不是对应的。
		if nums[i]+total[i-1] > total[i] {
			total = append(total, nums[i]+total[i-1])
		} else {
			total = append(total, total[i])
		}
	}

	return total[len(total)-1]
}
