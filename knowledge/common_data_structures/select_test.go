//  select_test.go
//
//
//  Created by JohnnyB0Y on 2021/07/05.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package commondatastructures_test

import (
	"fmt"
	"testing"
)

func SelectExample1() {
	c1 := make(chan int, 10)
	c2 := make(chan int, 10)
	c1 <- 1
	c2 <- 2

	select {
	case <-c1:
		fmt.Println("Run c1")
	case <-c2:
		fmt.Println("Run c2")
	}
}

func SelectExample2() {
	c := make(chan int)

	for i := 0; i < 3; i++ {
		go func() {
			fmt.Println("I'm Writing chan.........")
			c <- 11
		}()
	}

	for i := 0; i < 3; i++ {
		go func() {
			fmt.Println("I'm Reading chan.........")
			d := <-c
			fmt.Println(d)
		}()
	}

	fmt.Println("Before select chan!!!!")

	// 会阻塞
	for {
		select {
		case <-c:
			fmt.Println("Select reading chan ..........")
		case c <- 10:
			fmt.Println("Select writing chan ..........")
			// default:
			// 	fmt.Println("Select Default ..........")
		}
	}

	// 直接阻塞当前
	// select {}
}

func SelectExample3() {
	c := make(chan int, 10)
	c <- 10
	close(c)

	select {
	case d, ok := <-c:
		if !ok {
			fmt.Println("No data received!")
		}
		fmt.Println(d)
	}
}

func SelectExample4() {
	var c chan string

	select {
	case c <- "China":
		fmt.Println("Call send.")
	default:
		fmt.Println("Call default.")
	}
}

func TestSelect(t *testing.T) {
	// SelectExample1()
	// SelectExample2()
	// SelectExample3()
	// SelectExample4()

	// 初始化管道，并设置值为nil。称为零值。
	var c chan int
	fmt.Printf("chan type: %T, value: %v; nil type: %T, value: %v; chan == nil? %v \n", c, c, nil, nil, c == nil)
}

/**
运行测试命令
go test select_test.go -v -cover
*/

/**
select 语句，相当于 管道操作选择器；
当管道可操作时，会随机选择一个来执行；
当管道不可操作时，看有没有default选项，有就执行，没有就阻塞；

1）如果操作的管道为零值管道，由于零值管道不可读不可写，在运行时会被忽略掉；
2）从管道读取数据；
	- 当管道有数据可读时，当前case执行并返回；
	- 当管道没有数据可读，且管道已关闭，当前case执行并返回；
	- 当管道没有数据可读，且管道未关闭，当前case会被阻塞；
3）向管道写入数据；
	- 当管道的待读队列没有goroutine 排队时，且没有数据缓冲区 或 数据缓冲区已满时，当前case会被阻塞；
	- 当管道的待读队列没有goroutine 排队时，有数据缓冲区且未满时，当前case执行并返回；
	- 当管道的待读队列有goroutine 排队时，当前case执行并返回；
4）default 选项不会对管道进行操作。
	- 当管道不能读或不能写时，当前case执行并返回；
	- 当default case 存在时，select 语句永远不会被阻塞，一般在检测管道是否可用时会用到；

*/
