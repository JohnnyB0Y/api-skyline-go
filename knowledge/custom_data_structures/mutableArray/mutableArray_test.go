//  mutableArray_test.go
//
//
//  Created by JohnnyB0Y on 2021/07/10.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package mutableArray

import (
	"fmt"
	"testing"
)

func TestMutableArray(t *testing.T) {

	arrM := MakeMutableArray(1)
	arrM.Add(3)
	arrM.Add(4)
	arrM.Add(6)
	println(arrM.RemoveLast())
	// [  3   4 ]
	arrM.Print()
	arrM.Add(9)
	arrM.InsertAtIndex(10, 0)
	arrM.InsertAtIndex(11, 0)
	// [ 11  10   3   4   9 ]
	arrM.Print()
	arrM.InsertsFromIndex([]int{22, 23, 24}, 3)
	// [ 11  10   3  22  23  24   4   9 ]
	arrM.Print()

	arrM.Add(25)
	// [ 11  10   3  22  23  24   4   9  25 ]
	arrM.Print()

	arrM.Enumerate(func(i, v int) (stop bool) {
		// 打印前五元素
		fmt.Println(i, v)
		return i >= 4
	})

}
