//  string_test.go
//
//
//  Created by JohnnyB0Y on 2021/07/04.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package commondatastructures_test

import (
	"fmt"
	"strconv"
	"strings"
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

	// 6.0 字符串常用方法
	usedString()
}

func usedString() {
	// 统计字符数
	s1 := "Hello, 中国"
	r := []rune(s1)
	fmt.Println(r, len(r))
	for i := 0; i < len(r); i++ {
		fmt.Printf("%v ", r[i])
	}

	// 字符串转字节 []byte
	var s2 = []byte("Hello China!")
	fmt.Println(s2)

	// 字节转字符串
	s3 := string([]byte{97, 98, 99})
	fmt.Println(s3)

	// 10进制 转 2，8，16进制，返回对应的字符串
	str7 := strconv.FormatInt(123, 8)
	fmt.Println(str7)

	// 查找子串是否存在指定的字符串中
	contains := strings.Contains("Welcome!", "come")
	fmt.Println(contains)

	// 统计一个字符串有几个指定的字串
	count := strings.Count("Cheese", "e")
	fmt.Println(count)

	// 不区分大小写的字符串比较（==区分字母大小写）
	fold := strings.EqualFold("abc", "AbC")
	fmt.Println(fold)

	// 返回字串在字符串中第一次出现的索引
	index := strings.Index("China is a ...", "a")
	fmt.Println(index)

	// 返回字串在字符串最后一次出现的index，没有则返回-1
	lastIndex := strings.LastIndex("The google make go!", "go")
	fmt.Println(lastIndex)

	// 将指定的子串替换成另外一个字串
	replace := strings.Replace("I lean Chinese!", "lean", "study", 1)
	fmt.Println(replace)

	// 按照指定的某个字符，为分割标识，将一个字符串拆分为字符串数组
	split := strings.Split("I, am, the, only, one.", ",")
	fmt.Println(split)

	// 将字符串的字母大小写替换
	lower := strings.ToLower("Go to School.")
	upper := strings.ToUpper("Go to School.")
	fmt.Println(lower, upper)

	// 将字符串左右的空格去掉
	space := strings.TrimSpace(" Is me     ")
	fmt.Println(space)

	// 将字符串左右指定的字符去掉
	trim := strings.Trim("!Go!!", "!")
	fmt.Println(trim)

	// 将字符串左边/右边指定的字符去掉
	left := strings.TrimLeft("!bbq!", "!")
	right := strings.TrimRight("!bbq!", "!")
	fmt.Println(left, right)

	// 判断字符串以什么开头/结尾
	prefix := strings.HasPrefix("abcde", "ab")
	suffix := strings.HasSuffix("abcde", "de")
	fmt.Println(prefix, suffix)
}
