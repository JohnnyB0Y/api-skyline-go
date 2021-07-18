//  stack.go
//
//
//  Created by JohnnyB0Y on 2021/07/18.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package stack

import "errors"

type Stack struct {
	top int
	arr []string
}

func NewStack(size int) Stack {
	s := Stack{}
	s.top = -1
	s.arr = make([]string, size)
	return s
}

func (s *Stack) Empty() bool {
	return s.top == -1
}

func (s *Stack) Full() bool {
	return s.top+1 >= len(s.arr)
}

func (s *Stack) Push(item string) error {
	if s.Full() {
		return errors.New("is full baby")
	}

	s.top += 1
	s.arr[s.top] = item

	return nil
}

func (s *Stack) Pop() (item string, err error) {

	if s.Empty() {
		return item, errors.New("is empty baby")
	}

	item = s.arr[s.top]
	s.top -= 1

	return item, nil
}
