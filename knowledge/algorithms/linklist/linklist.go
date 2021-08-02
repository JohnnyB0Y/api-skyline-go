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
	if head == nil || head.Next == nil || k < 2 {
		return head
	}

	var hand *ListNode
	var tail *ListNode
	nodes := make([]*ListNode, 0, k)
	for head != nil {
		for i := 0; i < k; i++ {
			nodes = append(nodes, head)
			head = head.Next
			if head == nil {
				break
			}
		}
		if len(nodes) == k { // 满足 k个元素
			for i := k - 1; i > 0; i-- {
				if hand == nil {
					hand = nodes[i] // 第一个元素
				}
				nodes[i].Next = nodes[i-1] // 翻转元素
				nodes[i-1].Next = nil
			}
			if tail != nil {
				tail.Next = nodes[k-1] // 让链表尾部连上
			}
			tail = nodes[0]   // 记录链表尾部
			nodes = nodes[:0] // 清空数组
		} else { // 不满足k个元素
			if hand == nil { // 链表没有 k个元素
				return nodes[0]
			} else if len(nodes) > 0 { // 把最后的元素连接上
				tail.Next = nodes[0]
				return hand
			}
		}
	}
	return hand
}

func reverseKGroup2(head *ListNode, k int) *ListNode {
	var pre, cur, next, newHead, newTail *ListNode
	cur = head // 从第一个开始
	for cur != nil {
		for i := 0; i < k; i++ { // 少于 k个不翻转
			if cur == nil {
				if newTail != nil {
					newTail.Next = head // 可能是最后一组
					return newHead
				} else {
					return head
				}
			}
			cur = cur.Next
		}

		cur = head               // 交换开始
		for i := 0; i < k; i++ { // 翻转一组，共k个
			next = cur.Next // 记录下一个
			cur.Next = pre  // 翻转当前元素
			pre = cur       // 记录上一个
			cur = next      // 进入下一个
		}
		if newTail != nil {
			newTail.Next = pre // 尾巴指向当前组的头部
		}
		if newHead == nil {
			newHead = pre // 记录头部
		}
		newTail = head // 记录尾部
		head = cur     // 下一轮
		pre = nil
	}
	return newHead
}

/**
23. 合并K个升序链表
给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。

提示：

k == lists.length
0 <= k <= 10^4
0 <= lists[i].length <= 500
-10^4 <= lists[i][j] <= 10^4
lists[i] 按 升序 排列
lists[i].length 的总和不超过 10^4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-k-sorted-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) < 1 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	var head *ListNode
	var tail *ListNode
	smallers := make([]int, 0, 10) // 存放待插入的链表下标
	nilNodes := make([]int, 0, 10) // 记录空的链表下标

	for len(lists) > 0 {
		for i := 0; i < len(lists); i++ {
			if lists[i] == nil { // 记录nil链表
				nilNodes = append(nilNodes, i)
			} else if len(smallers) < 1 { // 记录首个
				smallers = append(smallers, i)
			} else { // 找最小
				if lists[i].Val < lists[smallers[0]].Val { // 小于就从0插入
					smallers = smallers[:1]
					smallers[0] = i
				} else if lists[i].Val == lists[smallers[0]].Val { // 等于就在后面拼接
					smallers = append(smallers, i)
				}
			}
		}

		// 拼接链表
		for i := 0; i < len(smallers); i++ {
			if head == nil {
				head = lists[smallers[i]] // 首个头部
			}
			if tail == nil {
				tail = head
			} else {
				tail.Next = lists[smallers[i]] // 指向下一个
				tail = lists[smallers[i]]
			}

			for tail.Next != nil {
				if tail.Next.Val != tail.Val { // 寻找相等的尾巴
					break
				}
				tail = tail.Next
			}

			if tail.Next != nil { // 更新 lists 链表
				lists[smallers[i]] = tail.Next
			} else { // 记录nil链表
				lists[smallers[i]] = nil
				nilNodes = append(nilNodes, smallers[i])
			}
		}
		smallers = smallers[:0]

		// 缩小空链表
		var last = len(lists) - 1
		for i := 0; i < len(nilNodes); i++ {
			if lists[nilNodes[i]] == nil {
				for last >= 0 {
					if lists[last] != nil { // 找到
						if nilNodes[i] < last { // 并且要小于 nilNodes[i] < last
							lists[nilNodes[i]] = lists[last]
							last--
						}
						break
					}
					last--
				}
			}
		}
		nilNodes = nilNodes[:0]
		lists = lists[:last+1]
	}
	return head
}
