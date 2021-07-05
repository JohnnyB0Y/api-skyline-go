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

	// 会阻塞
	select {
	case <-c:
		fmt.Println("Reading c.")
	case c <- 10:
		fmt.Println("Writing c.")
	}
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
	SelectExample1()
	// SelectExample2()
	SelectExample3()
	SelectExample4()
}

/**
运行测试命令
go test select_test.go -v -cover
*/
