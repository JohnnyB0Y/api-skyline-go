//  heap.go
//
//
//  Created by JohnnyB0Y on 2021/07/17.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package heap

import (
	"errors"
)

type heap struct {
	arr    []int
	length int
}

func NewHeap(arr []int) heap {
	var h heap = heap{}
	h.Build(arr)
	return h
}

func (h *heap) Pop() (int, error) {

	if h.length <= 0 {
		return 0, errors.New("no element baby")
	}

	// 最后的叶子 和 0 号位交换
	val := h.arr[0]
	h.arr[0] = h.arr[h.length-1]
	h.length -= 1
	h.heapify(0, true)
	return val, nil
}

func (h *heap) Push(val int) error {
	if h.length >= len(h.arr) {
		// 数组越界了
		return errors.New("out of bounds baby")
	}

	idx := h.length
	h.arr[idx] = val
	h.length += 1

	// left or right ?
	parent := idx
	for parent >= 0 {
		switch {
		case parent%2 != 0: // 左
			parent = (parent - 1) / 2
		case parent%2 == 0: // 右
			parent = (parent - 2) / 2
		}

		h.heapify(parent, false)
	}

	return nil
}

// 枚举遍历堆
func (h *heap) Enumeration(fn func(idx, val int) (stop bool)) {
	for i := 0; i < h.length; i++ {
		stop := fn(i, h.arr[i])
		if stop {
			break
		}
	}
}

func (h *heap) Build(arr []int) {
	h.arr = arr
	h.length = len(arr)
	// 最后一个父节点
	for i := h.length/2 - 1; i >= 0; i-- {
		h.heapify(i, false)
	}
}

func (h *heap) heapify(parent int, recursive bool) {

	if parent < 0 {
		return
	}

	// n / 2 = 叶子数
	if parent >= (h.length / 2) {
		return
	}

	// 大的节点上浮
	left := parent*2 + 1
	right := parent*2 + 2
	switch {
	case h.arr[left] > h.arr[right] && h.arr[left] > h.arr[parent]: // 左 > 右，且 左 > 父
		h.swap(left, parent)
		left = parent
	case h.arr[right] > h.arr[left] && h.arr[right] > h.arr[parent]: // 左 < 右，且 右 > 父
		h.swap(right, parent)
		right = parent
	}
	// 向下走
	if recursive {
		h.heapify(left, recursive)
		h.heapify(right, recursive)
	}
}

func (h *heap) swap(src, des int) {
	temp := h.arr[src]
	h.arr[src] = h.arr[des]
	h.arr[des] = temp
}
