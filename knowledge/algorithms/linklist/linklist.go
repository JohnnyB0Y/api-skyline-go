//  linklist.go
//
//
//  Created by JohnnyB0Y on 2021/07/31.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package linklist

import (
	"math/rand"
	"time"
)

/**
合并两个有序链表
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
*/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var head *ListNode
	var hand *ListNode = &ListNode{}
	for l1 != nil || l2 != nil {
		switch {
		case l1 == nil:
			hand.Next = l2
			l2 = l2.Next
		case l2 == nil:
			hand.Next = l1
			l1 = l1.Next
		case l1.Val >= l2.Val:
			if head == nil {
				head = l2
			}
			hand.Next = l2
			l2 = l2.Next
		case l1.Val < l2.Val:
			if head == nil {
				head = l1
			}
			hand.Next = l1
			l1 = l1.Next
		}
		hand = hand.Next
	}
	hand.Next = nil
	return head
}

/**
回文链表
请判断一个链表是否为回文链表。
https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnv1oc/
*/
func isPalindrome(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return true
	}

	nodes := make([]*ListNode, 0, 10)
	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}

	i, j := 0, len(nodes)-1
	for i < j {
		if nodes[i].Val != nodes[j].Val {
			return false
		}
		i++
		j--
	}

	return true
}

/**
环形链表
给定一个链表，判断链表中是否有环。

如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。

如果链表中存在环，则返回 true 。 否则，返回 false 。



进阶：

你能用 O(1)（即，常量）内存解决此问题吗？

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnwzei/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func hasCycle2(head *ListNode) bool {
	track := map[*ListNode]bool{}
	rand.Seed(time.Now().UnixNano())

	for head != nil {
		if node := track[head]; node { // 查水表
			return true
		}
		track[head] = true // 记录

		for i := 0; i < rand.Intn(6)+1; i++ { // 随机走几步
			if head != nil {
				head = head.Next
			} else {
				return false
			}
		}
	}
	return false
}

func hasCycle3(head *ListNode) bool {
	rand.Seed(time.Now().UnixNano())
	fast := head
	slow := head
	for fast != nil {
		step := rand.Intn(4) + 2
		for i := 0; i < step; i++ { // 随机走几步
			if fast != nil {
				fast = fast.Next
				if i%2 != 0 {
					slow = slow.Next
				}
				if fast == slow {
					return true
				}
			} else {
				return false
			}
		}
	}
	return false
}

func hasCycle(head *ListNode) bool {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

/**
25. K 个一组翻转链表
给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。

k 是一个正整数，它的值小于或等于链表的长度。

如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

进阶：
你可以设计一个只使用常数额外空间的算法来解决此问题吗？
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
*/

func reverseKGroup(head *ListNode, k int) *ListNode {

	return nil
}
