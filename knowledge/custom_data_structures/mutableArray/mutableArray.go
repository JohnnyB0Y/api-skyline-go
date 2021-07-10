//  mutableArray.go
//
//
//  Created by JohnnyB0Y on 2021/07/10.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

// 可变数组

package main

import "fmt"

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

// 添加，数组不够自动扩容
func (arrM *MutableArray) Add(value int) {
	index := arrM.end + 1
	if index < arrM.capacity {
		arrM.innerArr[index] = value
	} else {
		// 扩容两倍
		arrM.capacity *= 2
		newArr := heapArray(arrM.capacity)
		for i := arrM.start; i < arrM.length; i++ {
			newArr[i] = arrM.innerArr[i]
		}
		newArr[index] = value
		arrM.innerArr = newArr
	}

	arrM.end++
	arrM.length++
}

// 插入一个元素
func (arrM *MutableArray) InsertAtIndex(value int, index int) error {
	if index > arrM.length {
		// 返回插入越界错误❌
		return nil
	} else if index == arrM.length {
		arrM.Add(value)
		return nil
	}

	// 插入
	oldArr := arrM.innerArr
	if arrM.length >= arrM.capacity {
		// 扩容
		arrM.capacity *= 2
		newArr := heapArray(arrM.capacity)
		arrM.innerArr = newArr
		fmt.Print("ssssssssssssssssssssss\n")
		// 插入前半部分
		for i := arrM.start; i < index; i++ {
			fmt.Print(oldArr[i])
			arrM.innerArr[i] = oldArr[i]
		}
	}

	// 插入后半部分
	for i := arrM.length; i > index; i-- {
		arrM.innerArr[i] = oldArr[i-1]
	}

	// 插入当前
	arrM.innerArr[index] = value
	arrM.end++
	arrM.length++

	return nil
}

// 删除一个元素
func (arrM *MutableArray) RemoveAtIndex(index int) (int, error) {
	println("del", index)
	if index < arrM.length {
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
	// 错误❌处理放后面
	return 0, nil
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

func (arrM *MutableArray) Print() {
	fmt.Printf("MutableArray: [start: %d end: %d] length: %d, capacity: %d\n", arrM.start, arrM.end, arrM.length, arrM.capacity)
	for i := arrM.start; i < arrM.length; i++ {
		fmt.Println(i, " : ", arrM.innerArr[i])
	}
}

// 手动逃逸到堆上
func heapArray(capacity int) []int {
	s := make([]int, capacity)
	return s
}

func main() {
	arrM := MakeMutableArray(1)
	arrM.Add(3)
	arrM.Add(4)
	arrM.Add(6)
	println(arrM.RemoveLast())
	arrM.Add(9)
	arrM.InsertAtIndex(10, 0)
	arrM.InsertAtIndex(11, 0)
	arrM.Print()
}
