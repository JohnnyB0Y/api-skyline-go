//  array_test.go
//
//
//  Created by JohnnyB0Y on 2021/07/24.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package array

import (
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	t.Log(containsDuplicate([]int{1, 2, 3, 1}))
	t.Log(containsDuplicate([]int{1, 2, 3, 4}))
}

func TestIntersect(t *testing.T) {
	t.Log(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
	t.Log(intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
