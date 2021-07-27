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

/**
字符串中的第一个唯一字符
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
*/

func firstUniqChar(s string) int {
	if len(s) < 1 {
		return -1
	}
	charMap := map[byte]int{}

	for i := 0; i < len(s); i++ {
		charMap[s[i]] += 1
	}

	for i := 0; i < len(s); i++ {
		if charMap[s[i]] == 1 {
			return i
		}
	}
	return -1
}

// 不能给自带数据类型添加方法，So
type ModernString string

func NewModernString(s string) ModernString {
	return ModernString(s)
}

// 前面往后找第一个元素
func (s ModernString) FirstIndexOf(val byte) int {
	for i := 0; i < len(s); i++ {
		if s[i] == val {
			return i
		}
	}
	return -1
}

// 后面往前面找第一个元素
func (s ModernString) LastIndexOf(val byte) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == val {
			return i
		}
	}
	return -1
}

func firstUniqChar2(s string) int {
	ms := NewModernString(s)
	for i := 0; i < len(ms); i++ {
		if ms.FirstIndexOf(ms[i]) == ms.LastIndexOf(ms[i]) {
			return i
		}
	}
	return -1
}

/**
有效的字母异位词
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn96us/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	chars := make([]int, 26)
	for i := 0; i < len(s); i++ {
		chars[int(s[i])-97] += 1
		chars[int(t[i])-97] -= 1
	}

	for i := 0; i < len(chars); i++ {
		if chars[i] > 0 {
			return false
		}
	}

	return true
}

/**
验证回文串
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

说明：本题中，我们将空字符串定义为有效的回文串。
*/

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	var val1, val2 int8
	for left < right {
		val1 = int8(s[left])
		if !isCharAndNum(val1) { // 判断是否有效字母
			left++
			continue // 不符合要求，继续往下找
		}
		for {
			val2 = int8(s[right])
			if !isCharAndNum(val2) { // 判断是否有效字母
				right--
				continue // 不符合要求，继续往下找
			}
			if !isCharAndNumEqual(val1, val2) {
				return false // 不相等，不是回文
			}
			right-- // 是回文，进行下一轮
			break
		}
		left++ // 下一轮
	}
	return true
}

func isCharAndNumEqual(val1, val2 int8) bool {
	if val1 > 90 {
		val1 -= 32
	}
	if val2 > 90 {
		val2 -= 32
	}
	return val1 == val2
}

func isCharAndNum(val int8) bool {
	/**
	a-z：97-122
	A-Z：65-90
	0-9：48-57
	*/
	if val < 48 || val > 122 {
		return false
	}
	if val > 57 && val < 65 {
		return false
	}
	if val > 90 && val < 97 {
		return false
	}
	return true
}

/**
字符串转换整数 (atoi)
请你来实现一个 myAtoi(string s) 函数，使其能将字符串转换成一个 32 位有符号整数（类似 C/C++ 中的 atoi 函数）。

函数 myAtoi(string s) 的算法如下：

读入字符串并丢弃无用的前导空格
检查下一个字符（假设还未到字符末尾）为正还是负号，读取该字符（如果有）。 确定最终结果是负数还是正数。 如果两者都不存在，则假定结果为正。
读入下一个字符，直到到达下一个非数字字符或到达输入的结尾。字符串的其余部分将被忽略。
将前面步骤读入的这些数字转换为整数（即，"123" -> 123， "0032" -> 32）。如果没有读入数字，则整数为 0 。必要时更改符号（从步骤 2 开始）。
如果整数数超过 32 位有符号整数范围 [−231,  231 − 1] ，需要截断这个整数，使其保持在这个范围内。具体来说，小于 −231 的整数应该被固定为 −231 ，大于 231 − 1 的整数应该被固定为 231 − 1 。
返回整数作为最终结果。
注意：

本题中的空白字符只包括空格字符 ' ' 。
除前导空格或数字后的其余字符串外，请勿忽略 任何其他字符。


作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnoilh/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func myAtoi(s string) int {

}
