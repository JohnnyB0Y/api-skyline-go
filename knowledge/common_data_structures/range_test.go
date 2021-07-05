//  range_test.go
//
//
//  Created by JohnnyB0Y on 2021/07/05.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package commondatastructures_test

import (
	"fmt"
	"testing"
)

func TestRange(t *testing.T) {

	a := [...]string{"a", "b", "c"}
	// 作用于数组值，遍历数组
	for i, v := range a {
		fmt.Println(i, v)
	}

	// 作用于数组指针，遍历数组
	ap := &a
	for i, v := range ap {
		fmt.Println(i, v)
	}

	s := []int{1, 2, 3}
	// 作用于切片值，遍历切片
	for i, v := range s {
		fmt.Println(i, v)
	}

	// 不可作用于切片指针，遍历切片 ！！！！！
	sp := &s
	for i, v := range *sp {
		fmt.Println(i, v)
	}

}
