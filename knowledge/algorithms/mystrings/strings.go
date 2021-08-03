//  strings.go
//
//
//  Created by JohnnyB0Y on 2021/07/27.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package mystrings

import (
	"bytes"
	"strconv"
	"strings"
	"unsafe"
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
	// 取符号
	read := 0
	len_N := len(s)
	isDelete := false
	for read < len_N {
		if s[read] != ' ' && s[read] != '-' && s[read] != '+' {
			break // 不是空格，不是’-‘号
		} else if s[read] == '-' {
			isDelete = true
			read++
			break
		} else if s[read] == '+' {
			read++
			break
		}
		read++
	}
	// 取数字
	num := 0
	maxInt32 := 2147483648 // go 居然没有类型最值常量？？？
	for read < len_N {
		if s[read] <= '9' && s[read] >= '0' { // 有效数字
			num = num*10 + int(s[read]-'0')
			if num > maxInt32 {
				num = maxInt32
				break
			}
		} else { // 不是数字，退出
			break
		}
		read++
	}

	if isDelete {
		if num >= maxInt32 {
			return -maxInt32
		}
		return -num
	}
	if num >= maxInt32 {
		return maxInt32 - 1
	}
	return num
}

/**
实现 strStr()
实现 strStr() 函数。

给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回  -1 。



说明：

当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与 C 语言的 strstr() 以及 Java 的 indexOf() 定义相符。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnr003/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func strStr(haystack string, needle string) int {
	// 用这个 strings.Index(haystack, needle)，算不算不讲武德？
	// return strings.Index(haystack, needle)
	if len(needle) != 0 {
		matching := SubstringIndexsOf(haystack, needle)
		if len(matching) == 0 { // 没有匹配到
			return -1
		}
		return matching[0]
	}
	return 0 // 空字符串
}

func SubstringIndexsOf(haystack string, needle string) []int { // KMP 算法
	len_N := len(needle)
	len_H := len(haystack)
	matching := []int{}
	if len_N == 0 || len_H < len_N {
		return matching
	}

	ft := FallbackTableForStringNoMatching(needle)
	var i, j int
	for i < len_H {
		if haystack[i] != needle[j] { // 不匹配
			if j == 0 {
				i++ // needle的第一个元素比较失败，需要移动i
			} else {
				j = ft[j-1] // 回退
			}
			continue // 跳过
		}
		if j >= len_N-1 { // 匹配到一个
			matching = append(matching, i-len_N+1)
			if j == 0 {
				j = -1 // 单个元素匹配，为了抵消下面的 j++
			} else {
				j = ft[j-1]
			}
		}
		i++ // 如果还可以，下一轮
		j++
	}
	return matching
}

// 字符匹配失败的回退表
func FallbackTableForStringNoMatching(s string) map[int]int {
	ft := map[int]int{}
	if len(s) < 3 { // 字符数少于3，返回一个零值回退表
		return ft
	}
	i, j := 0, 1
	k, l, nextI := 0, 0, 0
	for j < len(s) {
		if s[i] == s[j] { // 匹配到
			ft[j] = i + 1
			i++
			j++
		} else { // 没有匹配到，下一轮
			nextI = 0
			if i > 0 && s[j] == s[i-1] { // a == a
				k = i - 2
				l = j - 1
				matching := true
				for k >= 0 { // 回叙匹配，极端情况下这里会消耗很多性能，应该有优化空间吧？
					if s[k] != s[l] {
						matching = false
						break
					}
					k--
					l--
				}
				if matching {
					ft[j] = ft[j-1]
					nextI = i // 下一个还会是 a 吗？
				}
			}
			i = nextI
			j++
		}
	}
	return ft
}

/**
外观数列
给定一个正整数 n ，输出外观数列的第 n 项。

「外观数列」是一个整数序列，从数字 1 开始，序列中的每一项都是对前一项的描述。

你可以将其视作是由递归公式定义的数字字符串序列：

countAndSay(1) = "1"
countAndSay(n) 是对 countAndSay(n-1) 的描述，然后转换成另一个数字字符串。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnpvdm/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func countAndSay(n int) string {
	nums := []byte{'1'}
	for i := 2; i < n+1; i++ {
		tmp := []byte{}
		read := nums[0]
		count := byte('1')
		for i := 1; i < len(nums); i++ {
			if read == nums[i] { // 相同
				count++
			} else { // 不同
				tmp = append(tmp, count, read)
				count = byte('1')
				read = nums[i]
			}
		}
		nums = append(tmp, count, read)
	}
	return *(*string)(unsafe.Pointer(&nums))
}

/**
最长公共前缀
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。
*/

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	// 我只想到了，字母遍历一个一个比较，遇到不同就返回；
	// 下面这个方法很巧妙！

	prefix := strs[0] // 把第一个字符串作为前缀
	for i := 0; i < len(strs); i++ {
		for strings.Index(strs[i], prefix) != 0 { // 判断字符串是否包含前缀，位置 0 开始才是前缀
			prefix = prefix[:len(prefix)-1]
		}
	}

	return prefix
}

/**
剑指 Offer II 096. 字符串交织
给定三个字符串 s1、s2、s3，请判断 s3 能不能由 s1 和 s2 交织（交错） 组成。

两个字符串 s 和 t 交织 的定义与过程如下，其中每个字符串都会被分割成若干 非空 子字符串：

s = s1 + s2 + ... + sn
t = t1 + t2 + ... + tm
|n - m| <= 1
交织 是 s1 + t1 + s2 + t2 + s3 + t3 + ... 或者 t1 + s1 + t2 + s2 + t3 + s3 + ...
提示：a + b 意味着字符串 a 和 b 连接。
0 <= s1.length, s2.length <= 100
0 <= s3.length <= 200
s1、s2、和 s3 都由小写英文字母组成
https://leetcode-cn.com/problems/IY6buf/
*/

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

}

// 没理解交错的意思！！！
func isInterleaveErrorVersion(s1 string, s2 string, s3 string) bool {

	if len(s1)+len(s2) != len(s3) {
		return false
	}

	chars := make([]uint8, 26)

	for i := 0; i < len(s3); i++ {
		idx := uint8(s3[i] - 'a')
		chars[idx] += 1
	}

	for i := 0; i < len(s1); i++ {
		idx := uint8(s1[i] - 'a')
		chars[idx] -= 1
	}
	for i := 0; i < len(s2); i++ {
		idx := uint8(s2[i] - 'a')
		chars[idx] -= 1
	}

	var count uint8
	for i := 0; i < 26; i++ {
		count += chars[i]
	}

	return count == 0
}
