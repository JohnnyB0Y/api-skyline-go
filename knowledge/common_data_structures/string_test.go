//  string_test.go
//
//
//  Created by JohnnyB0Y on 2021/07/04.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package commondatastructures_test

import (
	"fmt"
	"strconv"
	"testing"
)

func TestString(t *testing.T) {
	// 1.0 字符串长度，表示字符串的字节总数。
	s := "中国"
	fmt.Println(len(s)) // 6
	s = "🇨🇳"
	fmt.Println(len(s)) // 8
	s = "China"
	fmt.Println(len(s)) // 5

	// 2.0
	var s1 string
	s2 := "Hello World!"
	var s3 string
	s3 = "Hello s3!"

	fmt.Printf("s1 type: %T len: %d; s2 type: %T len: %d; s3 type: %T len: %d\n", s1, len(s1), s2, len(s2), s3, len(s3))

	// 3.0 双引号 与 反单引号
	s4 := "This is:\"MMU\"" // 双引号需要转义
	s5 := `This is:"MMU"`   // 反单引号不需要转义

	fmt.Println(s4, s5)

	// 4.0 字符串拼接
	s6 := "CC"
	s6 = s6 + "BB" + "AA"
	fmt.Println(s6)

	// 5.0 类型转换
	// https://mojotv.cn/go/golang-string-to-integer

	s7 := strconv.Itoa(1234567890)
	fmt.Println(s7)

	num, err := strconv.ParseInt(s7, 10, 64)
	if err == nil {
		fmt.Println(num)
	}
}
