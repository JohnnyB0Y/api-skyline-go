//  tree_test.go
//
//
//  Created by JohnnyB0Y on 2021/07/25.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package tree

import "testing"

func TestMaxDepth(t *testing.T) {
	root := TreeNode{}
	root.Val = 3

	root.Left = &TreeNode{}
	root.Left.Val = 9

	root.Right = &TreeNode{}
	root.Right.Val = 20

	root.Right.Left = &TreeNode{}
	root.Right.Left.Val = 15

	root.Right.Right = &TreeNode{}
	root.Right.Right.Val = 7

	t.Log(maxDepth(&root))
}

func TestIsSymmetric(t *testing.T) {
	root := TreeNode{}
	root.Val = 1

	root.Left = &TreeNode{}
	root.Left.Val = 2

	root.Right = &TreeNode{}
	root.Right.Val = 2

	root.Left.Right = &TreeNode{}
	root.Left.Right.Val = 3

	root.Right.Right = &TreeNode{}
	root.Right.Right.Val = 3

	t.Log(isSymmetric(&root))
}
