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

func removeDuplicates(nums []int) int {
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
