//  map_test.go
//
//
//  Created by JohnnyB0Y on 2021/07/04.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package commondatastructures_test

import (
	"fmt"
	"testing"
)

func InitMap() {
	// 字面量初始化
	m1 := map[string]int{
		"a": 1,
		"b": 2,
	}

	for _, v := range m1 {
		fmt.Println(v)
	}

	// make 初始化
	m2 := make(map[string]int, 10)
	m2["a"] = 1
	m2["b"] = 2
	m2["c"] = 3

	for k, v := range m2 {
		fmt.Println(k, ":", v)
	}
}

func TestMap(t *testing.T) {
	fmt.Println("Start map test -----------------------------")

	InitMap()

	fmt.Println("End map test -----------------------------")
}
