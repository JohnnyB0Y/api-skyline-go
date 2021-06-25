//  memory_test.go
//
//
//  Created by JohnnyB0Y on 2021/06/25.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package memory_test

import (
	"fmt"
	"runtime"
	"testing"
	"unsafe"
)

// memory layout
func bool_string_uint8_layout() string {
	type MyStruct struct {
		Field1 bool
		Field2 string
		Field3 uint8
	}

	var ms = MyStruct{}
	return fmt.Sprintf("totalSize: %d, %T : %d, %T %d, %T %d", unsafe.Sizeof(ms),
		ms.Field1, unsafe.Sizeof(ms.Field1),
		ms.Field2, unsafe.Sizeof(ms.Field2),
		ms.Field3, unsafe.Sizeof(ms.Field3))
}

func bool_int_uint8_layout() string {
	type MyStruct struct {
		Field1 bool
		Field2 int
		Field3 uint8
	}

	var ms = MyStruct{}
	return fmt.Sprintf("totalSize: %d, %T : %d, %T %d, %T %d", unsafe.Sizeof(ms),
		ms.Field1, unsafe.Sizeof(ms.Field1),
		ms.Field2, unsafe.Sizeof(ms.Field2),
		ms.Field3, unsafe.Sizeof(ms.Field3))
}

func bool_uint8_int_layout() string {
	type MyStruct struct {
		Field1 bool
		Field2 uint8
		Field3 int
	}

	var ms = MyStruct{}
	return fmt.Sprintf("totalSize: %d, %T : %d, %T %d, %T %d", unsafe.Sizeof(ms),
		ms.Field1, unsafe.Sizeof(ms.Field1),
		ms.Field2, unsafe.Sizeof(ms.Field2),
		ms.Field3, unsafe.Sizeof(ms.Field3))
}

func bool_uint8_uint8_layout() string {
	type MyStruct struct {
		Field1 bool
		Field2 uint8
		Field3 uint8
	}

	var ms = MyStruct{}
	return fmt.Sprintf("totalSize: %d, %T : %d, %T %d, %T %d", unsafe.Sizeof(ms),
		ms.Field1, unsafe.Sizeof(ms.Field1),
		ms.Field2, unsafe.Sizeof(ms.Field2),
		ms.Field3, unsafe.Sizeof(ms.Field3))
}

func bool_uint8_uint16_layout() string {
	type MyStruct struct {
		Field1 bool
		Field2 uint8
		Field3 uint16
	}

	var ms = MyStruct{}
	return fmt.Sprintf("totalSize: %d, %T : %d, %T %d, %T %d", unsafe.Sizeof(ms),
		ms.Field1, unsafe.Sizeof(ms.Field1),
		ms.Field2, unsafe.Sizeof(ms.Field2),
		ms.Field3, unsafe.Sizeof(ms.Field3))
}

func bool_uint16_uint8_layout() string {
	type MyStruct struct {
		Field1 bool
		Field2 uint16
		Field3 uint8
	}

	var ms = MyStruct{}
	return fmt.Sprintf("totalSize: %d, %T : %d, %T %d, %T %d", unsafe.Sizeof(ms),
		ms.Field1, unsafe.Sizeof(ms.Field1),
		ms.Field2, unsafe.Sizeof(ms.Field2),
		ms.Field3, unsafe.Sizeof(ms.Field3))
}

func bool_uint32_uint16_layout() string {
	// 匿名结构体
	var ms struct {
		Field1 bool
		Field2 uint32
		Field3 uint16
	}

	return fmt.Sprintf("totalSize: %d, %T : %d, %T %d, %T %d", unsafe.Sizeof(ms),
		ms.Field1, unsafe.Sizeof(ms.Field1),
		ms.Field2, unsafe.Sizeof(ms.Field2),
		ms.Field3, unsafe.Sizeof(ms.Field3))
}

func TestMemoryLayout(t *testing.T) {

	// totalSize: 32, bool : 1, string 16, uint8 1
	t.Log(bool_string_uint8_layout())

	// totalSize: 24, bool : 1, int 8, uint8 1
	t.Log(bool_int_uint8_layout())

	// totalSize: 16, bool : 1, uint8 1, int 8
	t.Log(bool_uint8_int_layout())

	// totalSize: 3, bool : 1, uint8 1, uint8 1
	t.Log(bool_uint8_uint8_layout())

	// totalSize: 4, bool : 1, uint8 1, uint16 2
	t.Log(bool_uint8_uint16_layout())

	// totalSize: 6, bool : 1, uint16 2, uint8 1
	t.Log(bool_uint16_uint8_layout())

	// totalSize: 12, bool : 1, uint32 4, uint16 2
	t.Log(bool_uint32_uint16_layout())

	/**
	Go 的内存布局和对齐
	对于结构体
	- 如果结构体中有字段类型的字节大小超过CPU字长
		- 那么，编译器会为 其余小于CPU字长的字段，填充字节，以对齐CPU字长，便于CPU访问数据；
	- 如果结构体中的最长的字段小于CPU字长
		- 那么，编译器会以其为标准，其余小于该字段字节大小的字段会填充字节，与该字段大小对齐；

	*/

	t.Log(printMyName(), printCallerName())
}

// 打印函数名
func printMyName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

// 打印函数调用者名字
func printCallerName() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}

/**
运行
go test -v -cover
*/
