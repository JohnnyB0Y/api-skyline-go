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
