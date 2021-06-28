package memory_test

import (
	"testing"
	"unsafe"
)

func TestEmptySlices(t *testing.T) {

	p1 := make([]int, 0)
	p2 := make([]int, 0)

	// 对比空切片是否指向同一块内存地址？
	t.Log(p1, p2)
	t.Logf("p1 arr: %p, p2 arr: %p", p1, p2)

	// 类型和数据结构大小
	t.Logf("p1 Type: %T size: %v, p2 Type: %T size: %v", p1, unsafe.Sizeof(p1), p2, unsafe.Sizeof(p2))

	// 切片的长度和容量
	t.Logf("p1 len: %d cap: %d, p2 len: %d cap: %d", len(p1), cap(p1), len(p2), cap(p2))

	// ----------------------------
	arr := [...]int{1, 2, 3}

	s1 := arr[0:0:0]
	s2 := arr[0:0:0]

	// 对比空切片是否指向同一块内存地址？
	t.Log(s1, s2)
	t.Logf("\n s1 pointer: %p arr pointer: %p,\n s2 pointer: %p arr pointer: %p", &s1, s1, &s2, s2)
	t.Logf("arr pointer: %p", &arr)

	// 类型和数据结构大小
	t.Logf("s1 Type: %T size: %v, s2 Type: %T size: %v", s1, unsafe.Sizeof(s1), s2, unsafe.Sizeof(s2))

	// 切片的长度和容量
	t.Logf("s1 len: %d cap: %d, s2 len: %d cap: %d", len(s1), cap(s1), len(s2), cap(s2))

}
