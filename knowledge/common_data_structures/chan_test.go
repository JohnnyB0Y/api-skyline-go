//  chan_test.go
//
//
//  Created by JohnnyB0Y on 2021/07/02.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package commondatastructures_test

import (
	"fmt"
	"testing"
)

func InitChan(t *testing.T) {
	{
		var ch1 chan int             // 声明管道，值为nil
		ch2 := make(chan int)        // 无缓冲管道
		ch3 := make(chan string, 10) // 带缓冲的管道，容量为10

		t.Log(ch1, ch2, ch3)
	}

	{
		ch := make(chan int, 10)
		ch <- 1   // 数据流入管道
		d := <-ch // 数据流出管道
		t.Log(d)
	}

	{
		ch := make(chan int, 10)
		chanParamReadAndWrite(ch)
		chanParamReadonly(ch)
		chanParamWriteonly(ch)
	}

}

// 管道函数
func chanParamReadAndWrite(ch chan int) {
	// 可读可写的管道
	ch <- 5
	d := <-ch
	fmt.Println(d)
}

func chanParamReadonly(ch <-chan int) {
	// 只读的管道
	d := <-ch
	fmt.Println(d)
}

func chanParamWriteonly(ch chan<- int) {
	// 只写的管道
	ch <- 8
	fmt.Println(ch)
}

func TestChan(t *testing.T) {

}
