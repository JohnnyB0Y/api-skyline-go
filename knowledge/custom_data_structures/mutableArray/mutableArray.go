//  mutableArray.go
//
//
//  Created by JohnnyB0Y on 2021/07/10.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

// 可变数组

package main

import (
	"errors"
	"fmt"
)

type MutableArray struct {
	innerArr []int
	length   int
	capacity int
	start    int
	end      int
}

func MakeMutableArray(capacity int) MutableArray {
	arrM := MutableArray{}
	arrM.innerArr = heapArray(capacity)
	arrM.start = 0
	arrM.end = -1
	arrM.length = 0
	arrM.capacity = capacity

	return arrM
}

func StandardMutableArray() MutableArray {
	return MakeMutableArray(24)
}

// 添加，数组不够自动扩容
func (arrM *MutableArray) Add(value int) error {
	return arrM.InsertAtIndex(value, arrM.end+1)
}

// 插入一个元素
func (arrM *MutableArray) InsertAtIndex(value int, index int) error {
	return arrM.InsertsFromIndex([]int{value}, index)
}

// 插入切片，自动扩容
func (arrM *MutableArray) InsertsFromIndex(values []int, index int) error {
	if index > arrM.length {
		// 返回插入越界错误❌
		return errors.New("数组下标越界")
	}

	size := len(values)

	// 插入
	oldArr := arrM.innerArr
	if arrM.length+size-1 >= arrM.capacity {
		// 扩容
		if size > arrM.capacity {
			arrM.capacity = size * 2
		} else {
			arrM.capacity = arrM.capacity * 2
		}
		arrM.innerArr = heapArray(arrM.capacity)
		// 插入前半部分
		for i := arrM.start; i < index; i++ {
			arrM.innerArr[i] = oldArr[i]
		}
	}

	// 插入后半部分
	for i := arrM.length + size - 1; i > index+size-1; i-- {
		arrM.innerArr[i] = oldArr[i-size]
	}

	// 插入当前
	for i := index; i < index+size; i++ {
		arrM.innerArr[i] = values[i-index]
	}

	arrM.end += size
	arrM.length += size

	return nil
}

// 删除一个元素
func (arrM *MutableArray) RemoveAtIndex(index int) (int, error) {
	if index >= arrM.start && index < arrM.length {
		item := arrM.innerArr[index]
		for i := index; i < arrM.length; i++ {
			if index+1 < arrM.length {
				arrM.innerArr[index] = arrM.innerArr[index+1]
				continue
			}
			// 最后一个置零
			arrM.innerArr[index] = 0
		}
		arrM.end--
		arrM.length--
		return item, nil
	}
	// 错误❌
	return 0, errors.New("数组下标越界")
}

// 尾元素
func (arrM *MutableArray) Last() (int, error) {
	return arrM.innerArr[arrM.end], nil
}

// 删除尾元素
func (arrM *MutableArray) RemoveLast() (int, error) {
	return arrM.RemoveAtIndex(arrM.end)
}

// 首元素
func (arrM *MutableArray) First() (int, error) {
	return arrM.innerArr[arrM.start], nil
}

// 删除首元素
func (arrM *MutableArray) RemoveFirst() (int, error) {
	return arrM.RemoveAtIndex(arrM.start)
}

// 获取元素
func (arrM *MutableArray) ValueAtIndex(index int) (int, error) {
	if index >= arrM.start && index < arrM.length {
		return arrM.innerArr[index], nil
	}
	return 0, errors.New("数组下标越界")
}

func (arrM *MutableArray) Print() {
	fmt.Printf("MutableArray: [start: %d end: %d] length: %d, capacity: %d\n", arrM.start, arrM.end, arrM.length, arrM.capacity)
	fmt.Print("[")
	for i := arrM.start; i < arrM.length; i++ {
		fmt.Printf("%3v ", arrM.innerArr[i])
	}
	fmt.Print("]\n")
}

// 手动逃逸到堆上
func heapArray(capacity int) []int {
	return make([]int, capacity)
}

func main() {
	arrM := MakeMutableArray(1)
	arrM.Add(3)
	arrM.Add(4)
	arrM.Add(6)
	println(arrM.RemoveLast())
	// [  3   4 ]
	arrM.Print()
	arrM.Add(9)
	arrM.InsertAtIndex(10, 0)
	arrM.InsertAtIndex(11, 0)
	// [ 11  10   3   4   9 ]
	arrM.Print()
	arrM.InsertsFromIndex([]int{22, 23, 24}, 3)
	// [ 11  10   3  22  23  24   4   9 ]
	arrM.Print()

	arrM.Add(25)
	// [ 11  10   3  22  23  24   4   9  25 ]
	arrM.Print()

}
