//  others.go
//
//
//  Created by JohnnyB0Y on 2021/07/30.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package others

/**
最小栈
设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

push(x) —— 将元素 x 推入栈中。
pop() —— 删除栈顶的元素。
top() —— 获取栈顶元素。
getMin() —— 检索栈中的最小元素。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnkq37/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

type MinStack struct {
	arr [][2]int
	top int
	min int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{[][2]int{}, -1, 0}
}

func (this *MinStack) Push(val int) {
	if this.top == -1 {
		this.min = val
	} else if this.arr[this.top][1] > val {
		this.min = val
	}

	this.top += 1
	if this.top < len(this.arr) {
		this.arr[this.top] = [2]int{val, this.min}
	} else {
		this.arr = append(this.arr, [2]int{val, this.min})
	}
}

func (this *MinStack) Pop() {
	this.arr = this.arr[:this.top]
	this.top -= 1
	if this.top != -1 {
		this.min = this.arr[this.top][1]
	}
}

func (this *MinStack) Top() int {
	return this.arr[this.top][0]
}

func (this *MinStack) GetMin() int {
	return this.arr[this.top][1]
}

/**
位1的个数
编写一个函数，输入是一个无符号整数（以二进制串的形式），返回其二进制表达式中数字位数为 '1' 的个数（也被称为汉明重量）。

提示：

请注意，在某些语言（如 Java）中，没有无符号整数类型。在这种情况下，输入和输出都将被指定为有符号整数类型，并且不应影响您的实现，因为无论整数是有符号的还是无符号的，其内部的二进制表示形式都是相同的。
在 Java 中，编译器使用二进制补码记法来表示有符号整数。因此，在上面的 示例 3 中，输入表示有符号整数 -3。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn1m0i/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func hammingWeight(num uint32) int {
	var count int8
	num2 := num << 1
	// 1110 1100 1000 0000 num2
	// 0111 0110 0100 0000 num2>>1
	// 1111 1110 1100 1000 num
	// +1    +1   +1   +1  = 4个1
	for num > 0 {
		if num2>>1 != num {
			count++
		}
		num = num << 1
		num2 = num2 << 1
	}
	return int(count)
}

/**
汉明距离
两个整数之间的 汉明距离 指的是这两个数字对应二进制位不同的位置的数目。

给你两个整数 x 和 y，计算并返回它们之间的汉明距离。
*/

func hammingDistance(x int, y int) int {
	var count int
	num := x ^ y
	for num > 0 {
		if num&1 == 1 {
			count++
		}
		num = num >> 1
	}
	return count
}

/**
颠倒二进制位
颠倒给定的 32 位无符号整数的二进制位。

提示：

请注意，在某些语言（如 Java）中，没有无符号整数类型。在这种情况下，输入和输出都将被指定为有符号整数类型，并且不应影响您的实现，因为无论整数是有符号的还是无符号的，其内部的二进制表示形式都是相同的。
在 Java 中，编译器使用二进制补码记法来表示有符号整数。因此，在上面的 示例 2 中，输入表示有符号整数 -3，输出表示有符号整数 -1073741825。

进阶:
如果多次调用这个函数，你将如何优化你的算法？

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnc5vg/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func reverseBits(num uint32) uint32 {
	var num2 uint32
	len_N := 32
	for len_N > 0 {
		num2 <<= 1
		num2 += num & 1
		num >>= 1
		len_N--
	}
	return num2
}

/**
杨辉三角
给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。

在「杨辉三角」中，每个数是它左上方和右上方的数的和。
*/
func generate(numRows int) [][]int {
	/**
	      [1]
	     [1,1]
	    [1,2,1]
	   [1,3,3,1]
	  [1,4,6,4,1]
	*/
	results := make([][]int, 0, numRows)
	switch {
	case numRows > 0 && numRows < 2:
		results = append(results, []int{1})
	case numRows >= 2:
		results = append(results, []int{1}, []int{1, 1})
	}
	for i := 2; i < numRows; i++ {
		result := make([]int, 0, i+1)
		result = append(result, 1) // 前

		lastResult := results[i-1]
		for k := 0; k < i-1; k++ { // 中
			result = append(result, lastResult[k]+lastResult[k+1])
		}

		result = append(result, 1) // 后
		// 放入
		results = append(results, result)
	}
	return results
}

/**
有效的括号
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnbcaj/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func isValid(s string) bool {

	stack := make([]byte, 0, 10)
	for i := 0; i < len(s); i++ {
		if isRight(s[i]) { // 右边，需要从栈中取出配对的括号
			len_N := len(stack)
			if len_N < 1 || !isPair(stack[len_N-1], s[i]) { // 栈里没元素，说明没有配对
				return false
			}
			stack = stack[:len_N-1] // 如果配对,出栈
		} else {
			stack = append(stack, s[i]) // 左括号往里塞
		}
	}
	return len(stack) == 0
}

func isPair(left, right byte) bool {
	switch {
	case left == '(' && right == ')':
		return true
	case left == '[' && right == ']':
		return true
	case left == '{' && right == '}':
		return true
	}
	return false
}

func isRight(left byte) bool {
	switch {
	case left == ')':
		return true
	case left == ']':
		return true
	case left == '}':
		return true
	}
	return false
}

/**
缺失数字
给定一个包含 [0, n] 中 n 个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数。



进阶：

你能否实现线性时间复杂度、仅使用额外常数空间的算法解决此问题?


作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnj4mt/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func missingNumber(nums []int) int {
	len_N := len(nums)
	if len_N == 1 {
		if nums[0] == 0 {
			return 1
		} else {
			return 0
		}
	}

	total := (1 + len_N) * len_N / 2
	for i := 0; i < len_N; i++ {
		total -= nums[i]
	}
	return total
}
