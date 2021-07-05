//  range_test.go
//
//
//  Created by JohnnyB0Y on 2021/07/05.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package commondatastructures_test

import (
	"fmt"
	"testing"
	"time"
)

func TestRange(t *testing.T) {

	a := [...]string{"b", "c"}
	// 作用于数组值，遍历数组
	for i, v := range a {
		fmt.Println(i, v)
	}

	// 作用于数组指针，遍历数组
	ap := &a
	for i, v := range ap {
		fmt.Println(i, v)
	}

	s := []int{1, 2}
	// 作用于切片值，遍历切片
	for i, v := range s {
		fmt.Println(i, v)
	}

	// 不可作用于切片指针，遍历切片 ！！！！！
	sp := &s
	for i, v := range *sp {
		fmt.Println(i, v)
	}

	stringDemo()
	mapDemo()
	chanDemo()

}

func stringDemo() {
	str := "Hi 米粉"
	// 作用于字符串值，遍历字符串
	for i, v := range str {
		fmt.Printf("%d %c\n", i, v)
	}

	// 不可作用于字符串指针，遍历字符串
	strp := &str
	for i, v := range *strp {
		fmt.Printf("%d %c\n", i, v)
	}
}

func mapDemo() {
	m := map[string]string{"a": "apple", "b": "boy", "c": "carter"}
	// 作用于映射集合值，遍历映射集合
	for k, v := range m {
		fmt.Printf("%s for %s\n", k, v)
	}

	// 不可作用于映射集合指针，遍历映射集合
	mp := &m
	for k, v := range *mp {
		fmt.Printf("%s for %s\n", k, v)
	}
}

func chanDemo() {
	c := make(chan string, 2)
	c <- "Kobe"
	c <- "Bryant"

	// 定时
	time.AfterFunc(time.Microsecond, func() {
		// 关闭管道
		fmt.Printf("Before close the channel.\n")
		close(c)
		fmt.Printf("After close the channel.\n")
	})

	// 这里遍历完两个元素后会阻塞，直到定时器关闭channel后才结束遍历。
	for e := range c {
		fmt.Printf("%s\n", e)
	}
}
