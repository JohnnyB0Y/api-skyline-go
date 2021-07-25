//  tree.go
//
//
//  Created by JohnnyB0Y on 2021/07/25.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package tree

import (
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
