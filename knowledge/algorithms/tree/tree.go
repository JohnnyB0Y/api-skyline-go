//  tree.go
//
//
//  Created by JohnnyB0Y on 2021/07/25.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package tree

import (
	"fmt"
	"math"
)

/*
二叉树的最大深度
给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。
*/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	if left > right {
		return left + 1
	}
	return right + 1
}

/**
验证二叉搜索树
给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：

节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn08xg/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func isValidBST(root *TreeNode) bool {
	return isValidBSTInScope(root, math.MinInt64, math.MaxInt64)
}

func isValidBSTInScope(root *TreeNode, leftMin, rightMax int) bool {
	if root == nil {
		return true
	}

	// 左 > root ? 并且 左 < min ？
	if root.Left != nil && (root.Left.Val >= root.Val || root.Left.Val <= leftMin) {
		return false
	}

	// 右 < root ?
	if root.Right != nil && (root.Right.Val <= root.Val || root.Right.Val >= rightMax) {
		return false
	}

	// 左右 == true ？
	return isValidBSTInScope(root.Left, leftMin, root.Val) && isValidBSTInScope(root.Right, root.Val, rightMax)
}

/**
对称二叉树
给定一个二叉树，检查它是否是镜像对称的。


例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3


但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn7ihv/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	// 不对称
	if (root.Left == nil && root.Right != nil) || (root.Left != nil && root.Right == nil) {
		return false
	}

	// 左子树 == 右子树
	left := sumOfTree(root.Left, true)
	right := sumOfTree(root.Right, false)
	fmt.Println(left, right)
	return left == right
}

// 计算树的总值
func sumOfTree(root *TreeNode, isLeft bool) int {
	if root == nil {
		return 0
	}

	left := sumOfTree(root.Left, isLeft)
	right := sumOfTree(root.Right, isLeft)

	if isLeft { // * 10 * 3 是为了扩大数据的差异，便于比较差异，实际是有漏洞的
		return (right*10 - left*3) + root.Val
	}

	return (left*10 - right*3) + root.Val
}

/**
将有序数组转换为二叉搜索树
给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 高度平衡 二叉搜索树。

高度平衡 二叉树是一棵满足「每个节点的左右两个子树的高度差的绝对值不超过 1 」的二叉树。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xninbt/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func sortedArrayToBST(nums []int) *TreeNode {

	if len(nums) < 1 {
		return nil
	}

	// 中间元素
	pivot, leftArr, rightArr := slicingSortedArray(nums)

	// 左右
	pivot.Left = sortedArrayToBST(leftArr)
	pivot.Right = sortedArrayToBST(rightArr)

	return pivot
}

// 切割数组
func slicingSortedArray(nums []int) (pivot *TreeNode, leftArr, rightArr []int) {

	if len(nums) < 1 {
		return
	}

	if len(nums) == 1 {
		pivot = &TreeNode{}
		pivot.Val = nums[0]
		return
	}

	mid := len(nums) / 2
	pivot = &TreeNode{}
	pivot.Val = nums[mid]
	if mid != 0 {
		leftArr = nums[0:mid]
	}
	if mid+1 < len(nums) {
		rightArr = nums[mid+1:]
	}
	return
}

/**
二叉树的层序遍历
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnldjj/
*/
type treesNodes = [][]*TreeNode
type treesNode = []*TreeNode

func levelOrder(root *TreeNode) [][]int {

	if root == nil {
		return nil
	}

	q := treesNodes{
		treesNode{root.Left, root.Right},
	}
	arr := [][]int{{root.Val}}

	for len(q) > 0 {
		// 出队
		nodes := q[0]
		q = q[1:]

		vals := []int{}
		children := treesNode{}

		for i := 0; i < len(nodes); i++ {
			node := nodes[i]
			if node != nil {
				// 放入数据
				vals = append(vals, node.Val)
				children = append(children, node.Left, node.Right)
			}
		}

		if len(vals) > 0 {
			arr = append(arr, vals)
		}

		if len(children) > 0 {
			q = append(q, children)
		}
	}

	return arr
}
