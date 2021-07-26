//  strings.go
//
//
//  Created by JohnnyB0Y on 2021/07/27.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package mystrings

import (
	"bytes"
	"strconv"
)

/**
反转字符串
编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 char[] 的形式给出。

不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。

你可以假设数组中的所有字符都是 ASCII 码表中的可打印字符。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnhbqj/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func reverseString(s []byte) {
	if s == nil || len(s) < 2 {
		return
	}

	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

/**
整数反转
给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。

如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。

假设环境不允许存储 64 位整数（有符号或无符号）。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnx13t/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func reverse(x int) int {

	xStr := strconv.Itoa(x)
	isMinus := false
	if x < 0 {
		isMinus = true
		xStr = xStr[1:]
	}

	var buffer bytes.Buffer
	for i := len(xStr) - 1; i >= 0; i-- {
		buffer.WriteByte(xStr[i])
	}

	var newX64, err = strconv.ParseInt(buffer.String(), 10, 64)
	if err != nil {
		return 0
	}

	if newX64 != int64(int32(newX64)) { // 不接受 64 位
		return 0
	}

	if isMinus {
		return -int(newX64)
	}

	return int(newX64)
}

func reverse2(x int) int {
	var x64 int64 = int64(x)
	var newX64 int64 = 0
	// newX64 = 0   x64 = 123 (初始状态)
	// newX64 = 3   x64 = 12
	// newX64 = 32  x64 = 1
	// newX64 = 321 x64 = 0
	for x64 != 0 {
		newX64 = newX64*10 + (x64 % 10) // 取出最后一位数
		x64 /= 10                       // 去掉最后一位数
	}

	if newX64 != int64(int32(newX64)) { // 不接受 64 位
		return 0
	}
	return int(newX64)
}
