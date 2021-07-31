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
