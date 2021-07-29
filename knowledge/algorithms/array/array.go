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

import "fmt"

func RemoveDuplicates(nums []int) int {
	N := len(nums)
	if N < 2 {
		return N
	}

	cur, next := 0, 1
	for next < N {
		// 不相等
		if nums[cur] != nums[next] {
			for i := 1; i < N-next+1; i++ {
				nums[cur+i] = nums[next+i-1]
			}
			N -= next - cur - 1
			cur++
			next = cur + 1
			continue
		}
		// 相等，继续向前探
		next++

		// next走到最后 ? 把后面的重复项去除
		if next >= N {
			for i := 1; i < N-next+1; i++ {
				nums[cur+i] = nums[next+i-1]
			}
			N -= next - cur - 1
		}
	}
	return N
}

func RemoveDuplicates2(nums []int) int {
	N := len(nums)
	if N < 2 {
		return N
	}
	// fill 待填充的数组下标, find 寻找不同数字的数组下标
	fill, find := 1, 1
	for find < N {
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
	N := len(prices)
	if N < 2 {
		return 0
	}
	//  0 1 2 3 4 5
	// [7,1,5,3,6,4]
	buy, sale, profit := 0, 0, 0
	for buy < N-1 {
		// 当前价格高于下一个价格
		if prices[buy] > prices[buy+1] {
			buy++ // 继续找买入点
			continue
		}

		// 走到这里，算是找到了买入点，开始找卖出点
		sale = buy + 1
		for sale < N {
			if sale == N-1 || prices[sale] > prices[sale+1] { // sale 是最后一个 或 比下一个价格高 就不找了
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

func Rotate(nums []int, k int) {
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

	// k = 3
	// [1,2,3,4,5,6]
	// 1 -> 4; 4 -> 1;
	// 2 -> 5; 5 -> 2;
	// 3 -> 6; 6 -> 3;
	// [4,5,6,1,2,3]

	// 倒着走
	N := len(nums)

	if N < k { // k 大于 数组个数时，去掉多余移动的步数
		k = k % N
	}

	if k == 0 || N < 1 || N == k {
		return
	}

	// 临时存放 待覆盖 计数   轮数
	var temp, fill, count, round int
	for count < N {
		fill = N - k + round
		temp = nums[round]
		nums[round] = nums[fill]
		count += 1
		for fill != round {
			next := fill - k
			if next < 0 { // 循环下标
				next = N + next
			}

			if next == round { // 回到原点，此轮结束！
				nums[fill] = temp
				count += 1
				break
			}

			nums[fill] = nums[next]
			count += 1
			fill = next
		}
		round += 1 // 下一轮
	}
}

/**
存在重复元素
给定一个整数数组，判断是否存在重复元素。

如果存在一值在数组中出现至少两次，函数返回 true 。如果数组中每个元素都不相同，则返回 false 。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x248f5/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func containsDuplicate(nums []int) bool {
	N := len(nums)
	if N <= 1 {
		return false
	}

	set := make(map[int]bool)
	for i := 0; i < N; i++ {
		if _, ok := set[nums[i]]; ok {
			// 找到
			return true
		}
		set[nums[i]] = true
	}

	return false
}

/**
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

说明：

你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x21ib6/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func singleNumber(nums []int) int {

	// 看了下答案，用异或操作 😁
	if len(nums) == 1 {
		return nums[0]
	}

	var i = 2
	var val int = nums[0] ^ nums[1]
	for i < len(nums) {
		val = val ^ nums[i]
		i++
	}
	return val
}

/**
两个数组的交集 II
给定两个数组，编写一个函数来计算它们的交集。
*/

func intersect(nums1 []int, nums2 []int) []int {
	src := nums1
	des := nums2
	if len(nums1) > len(nums2) {
		src = nums2
		des = nums1
	}

	// nums1 = [1,2,2,1], nums2 = [2,2]
	set := make(map[int]int)
	for i := 0; i < len(src); i++ {
		set[src[i]] += 1
	}

	result := make([]int, 0, 4)
	for i := 0; i < len(des); i++ {
		if val := set[des[i]]; val > 0 {
			// 找到了
			result = append(result, des[i])
			set[des[i]] -= 1
		}
	}
	return result
}

/**
加一
给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2cv1c/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func plusOne(digits []int) []int {
	if digits[0] == 0 {
		return []int{1}
	}

	N := len(digits)
	if digits[N-1] != 9 {
		digits[N-1] += 1
		return digits
	}

	for i := N - 1; i >= 0; i-- {
		digits[i] = digits[i] + 1
		if digits[i]%10 == 0 {
			digits[i] = 0
		} else {
			break
		}
	}

	if digits[0] == 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}

/**
移动零
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
*/

func moveZeroes(nums []int) {
	start := 0
	for i := 1; i < len(nums); i++ {
		if nums[start] == 0 && nums[i] != 0 { // 不等于 0，交换
			nums[start], nums[i] = nums[i], nums[start]
		}
		if nums[start] != 0 { // start 不等于0 跟着走
			start++
		}
	}
}

/**
两数之和
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
你可以按任意顺序返回答案。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2jrse/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func twoSum2(nums []int, target int) []int {
	appear := map[int]int{}
	for i := 0; i < len(nums); i++ {
		remainder := target - nums[i]         // target - nums[i] = remainder；
		if idx, ok := appear[remainder]; ok { // remainder 出现过 ？
			return []int{idx, i} // target = nums[i] + nums[idx]
		}
		appear[nums[i]] = i // 未命中，在 appear 记录
	}
	return []int{}
}

/**
有效的数独
请你判断一个 9x9 的数独是否有效。只需要 根据以下规则 ，验证已经填入的数字是否有效即可。

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）
数独部分空格内已填入了数字，空白格用 '.' 表示。

注意：

一个有效的数独（部分已被填充）不一定是可解的。
只需要根据以上规则，验证已经填入的数字是否有效即可。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2f9gg/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func isValidSudoku2(board [][]byte) bool {
	appear := map[int]bool{}
	len_Row := len(board)
	len_Col := len_Row

	for row := 0; row < len_Row; row++ {
		for col := 0; col < len_Col; col++ {
			if board[row][col] == '.' {
				continue // 这个略过
			}
			// row 行字典
			keyR := 10000 + row*100 + int(board[row][col]) // 0 ~ 9 行
			// col 列字典
			keyC := 20000 + col*100 + int(board[row][col]) // 0 ~ 9 列
			// 3x3网格字典
			keyM := 30000 + row/3*100 + col/3*1000 + int(board[row][col]) // 0 ~ 9 3x3网格

			if has := appear[keyR]; has { // 在里面找到了，就无效了
				return false
			}
			appear[keyR] = true // 把它加进去

			if has := appear[keyC]; has { // 在里面找到了，就无效了
				return false
			}
			appear[keyC] = true // 把它加进去

			if has := appear[keyM]; has { // 在里面找到了，就无效了
				return false
			}
			appear[keyM] = true // 把它加进去
		}
	}
	return true
}

func isValidSudoku(board [][]byte) bool {
	// 答案提到用位运算！！！！
	// 位运算 49 ~ 57
	appearRow := make([]int16, 9)
	appearCol := make([]int16, 9)
	appearMesh := make([]int16, 9)
	int16_1 := int16(1)

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == '.' {
				continue // 这个略过
			}
			target := int16(board[row][col] - '0') // 算出数字 0 ~ 9
			bitIdx := int16_1 << target
			fmt.Println(target, ":", bitIdx)
			// 在 row 里面找？？？
			fmt.Println("row:", appearRow[row]&bitIdx > 0)
			if appearRow[row]&bitIdx > 0 { // 与运算，1-1 得 1；1-0，0-0 得 0；
				return false
			}
			appearRow[row] |= bitIdx // 或运算，1-0，1-1 得 1；0-0 得 0；

			// 在 col 里面找？？？
			fmt.Println("col:", appearCol[col]&bitIdx > 0)
			if appearCol[col]&bitIdx > 0 {
				return false
			}
			appearCol[col] |= bitIdx // 把它加进去

			// 在网格找？？？
			// 0, 1，2，3，4，5，6，7，8
			meshIdx := (row/3)*3 + col/3
			fmt.Println("mesh:", appearMesh[meshIdx]&bitIdx > 0)
			if appearMesh[meshIdx]&bitIdx > 0 { // 在里面找到了，就无效了
				return false
			}
			appearMesh[meshIdx] |= bitIdx // 把它加进去

		}
	}
	return true
}

/**
旋转图像
给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。

你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnhhkv/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func rotate(matrix [][]int) {

	len_N := len(matrix)
	if len_N < 2 {
		return
	}
	// i 走 i - 1 圈
	var tmp int
	for i := 0; i < len_N-1; i++ {
		tmp = matrix[i][0]
		// 1 4

		// (3,0) ==> (0,0)
		// (3,1) ==> (1,0)
		// (3,2) ==> (2,0)
		matrix[i][0] = matrix[len_N-1][i]
		// 7 -> 1; 8 -> 4

		// (3,3) ==> (3,0)
		// (2,3) ==> (3,1)
		// (1,3) ==> (3,2)
		matrix[len_N-1][i] = matrix[len_N-1-i][len_N-1]
		// 9 -> 7

		// (0,3) ==> (3,3)
		// (0,2) ==> (2,3)
		// (0,1) ==> (1,3)
		matrix[len_N-1-i][len_N-1] = matrix[0][len_N-1-i]
		// 3 -> 9

		// (0,0) ==> (0,3)
		// (1,0) ==> (0,2)
		// (2,0) ==> (0,1)
		matrix[0][len_N-1-i] = tmp
		// 1 -> 3
	}

	if len_N-2 < 2 {
		return
	}

	subMatrix := make([][]int, len_N-2)
	for i := 1; i < len_N-1; i++ {
		subMatrix[i-1] = matrix[i][1 : len_N-1]
	}
	// 继续
	rotate(subMatrix)
}

/**
合并两个有序数组
给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。

初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。你可以假设 nums1 的空间大小等于 m + n，这样它就有足够的空间保存来自 nums2 的元素。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnumcr/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	len_Nk := len(nums1)      // 完整nums1 数组长度
	len_Nj := len(nums2)      // nums2 数组长度
	len_Ni := len_Nk - len_Nj // nums1 有效数字的数组长度
	i := len_Ni - 1
	j := len_Nj - 1
	k := len_Nk - 1
	for k >= 0 {
		switch {
		case i < 0: // i 走完
			nums1[k] = nums2[j]
			k--
			j--
		case j < 0: // j 走完
			nums1[k] = nums1[i]
			k--
			i--
		case nums1[i] >= nums2[j]:
			nums1[k] = nums1[i]
			k--
			i--
		case nums1[i] < nums2[j]:
			nums1[k] = nums2[j]
			k--
			j--
		}
	}
}

/**
第一个错误的版本
你是产品经理，目前正在带领一个团队开发新的产品。不幸的是，你的产品的最新版本没有通过质量检测。由于每个版本都是基于之前的版本开发的，所以错误的版本之后的所有版本都是错的。

假设你有 n 个版本 [1, 2, ..., n]，你想找出导致之后所有版本出错的第一个错误的版本。

你可以通过调用 bool isBadVersion(version) 接口来判断版本号 version 是否在单元测试中出错。实现一个函数来查找第一个错误的版本。你应该尽量减少对调用 API 的次数。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnto1s/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func firstBadVersion(n int) int {
	if n == 1 {
		return n
	}

	// 1 是不是bad
	if isBadVersion(1) {
		return 1
	}

	left := 1
	right := n
	pivot := n / 2
	// 5 4
	for left < right-1 {
		// pivot = 2, 3, 4
		if isBadVersion(pivot) {
			// right = 4
			right = pivot
			pivot = (pivot-left)/2 + left
		} else {
			// left = 2, 3
			left = pivot
			pivot = (right-pivot)/2 + pivot
		}
	}
	return right
}

// 模拟而已
func isBadVersion(version int) bool {
	return true
}
