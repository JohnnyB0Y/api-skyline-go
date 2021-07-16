package main

import (
	"errors"
	"fmt"
)

func main() {
	if err := doit(true); err != nil {
		// 出错直接返回
		fmt.Println(err)
		return
	}

	if err := doit(false); err != nil {
		// 出错直接返回
		fmt.Println(err)
		return
	}

	fmt.Println("The end.")
}

func doit(it bool) error {
	if it {
		return nil
	}
	return errors.New("error")
}
