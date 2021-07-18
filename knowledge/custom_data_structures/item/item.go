//  item.go
//
//
//  Created by JohnnyB0Y on 2021/07/18.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package item

type Item struct {
	Key *interface{}
	Val *interface{}
}

// const (
// 	// 相等
// 	equal = iota
// 	// 不相等
// 	notEqual
// 	// 小于
// 	less
// 	// 小于等于
// 	lessOrEqual
// 	// 大于
// 	greater
// 	// 大于等于
// 	greaterOrEqual
// )

// 比较 i1 和 i2的key；返回 是否 i1 == i2 ？，i1 < i2 ？，i1 > i2 ？
func Compare(i1, i2 Item) (isEqual, isLess, isGreater bool) {
	switch (*i1.Key).(type) {
	case int, int8, int16, int32, int64:
		k1 := (*i1.Key).(int64)
		k2 := (*i2.Key).(int64)
		isEqual = k1 == k2
		isGreater = k1 > k2
		isLess = k1 < k2
	case uint, uint8, uint16, uint32, uint64:
		k1 := (*i1.Key).(uint64)
		k2 := (*i2.Key).(uint64)
		isEqual = k1 == k2
		isGreater = k1 > k2
		isLess = k1 < k2
	case string:
		k1 := (*i1.Key).(string)
		k2 := (*i2.Key).(string)
		isEqual = k1 == k2
		isGreater = k1 > k2
		isLess = k1 < k2
	}
	return
}
