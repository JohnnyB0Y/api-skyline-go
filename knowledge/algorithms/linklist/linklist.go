//  linklist.go
//
//
//  Created by JohnnyB0Y on 2021/07/31.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package linklist

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
