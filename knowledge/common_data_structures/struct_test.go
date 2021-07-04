//  struct_test.go
//
//
//  Created by JohnnyB0Y on 2021/07/04.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package commondatastructures_test

import (
	"fmt"
	"reflect"
	"testing"
)

type Animal struct {
	Name string
}

func (a *Animal) SetName(name string) {
	a.Name = name
}

type Cat struct {
	Animal
	Age uint8
}

type Dog struct {
	a   Animal
	Age uint8
}

func InitStruct() {
	cat := Cat{}

	// 各种赋值访问方式
	cat.SetName("(●—●)")
	fmt.Println(cat.Name)

	cat.Name = "小白"
	fmt.Println(cat.Name)

	cat.Animal.Name = "小黑"
	fmt.Println(cat.Name)

	cat.Animal.SetName("肥猫")
	fmt.Println(cat.Name)

	dog := Dog{}

	dog.a.Name = "大狗"
	fmt.Println(dog.a.Name)

	dog.a.SetName("小狗")
	fmt.Println(dog.a.Name)

	dog.a.SetName("肥狗")
	fmt.Println(dog.a.Name)

	// dog.Name undefined (type Dog has no field or method Name)compiler
	// dog.Name = "不可访问"

	// 一些初始化方法
	cat2 := Cat{
		Animal{
			"小🐈",
		},
		13,
	}

	dog2 := Dog{Age: 12, a: Animal{
		"小dog",
	}}
	fmt.Println(cat2, dog2)
}

type CustomType struct {
	Field string `json:"kind,omitempty" protobuf:"bytes,1,opt,name=kind"`
}

// Tag 结构体标记，相当于类属性的描述符吧
// 举例： Field string `json:"kind,omitempty" protobuf:"bytes,1,opt,name=kind"`
// Field 字段中的Tag 包含两个 key:"value" 对，分别是 json:"kind,omitempty" 和 protobuf:"bytes,1,opt,name=kind"
// key 一般表示用途，例如 json 表示用于控制结构体类型与 JSON 格式数据之间的转化，protobuf 表示序列化和反序列化控制。
// value 一般表示控制指令，具体指令的行为由对应实现的库实现。

func TagStruct() {
	// 通过反射，回去tag信息
	custom := CustomType{}
	ct := reflect.TypeOf(custom)

	for i := 0; i < ct.NumField(); i++ {
		json := ct.Field(i).Tag.Get("json")
		protobuf := ct.Field(i).Tag.Get("protobuf")
		fmt.Printf("Field: %s, Tag: json: %s, protobuf: %s\n", ct.Field(i).Name, json, protobuf)
	}

	// 通过Tag信息结合反射，可以对结构体进行配置处理。
}

func TestStruct(t *testing.T) {
	InitStruct()
	TagStruct()
}
