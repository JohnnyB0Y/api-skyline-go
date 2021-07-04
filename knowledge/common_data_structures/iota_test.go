//  iota_test.go
//
//
//  Created by JohnnyB0Y on 2021/07/04.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package commondatastructures_test

import (
	"fmt"
	"testing"
)

const (
	bit0, mask0 = 1 << iota, 1<<iota - 1
	bit1, mask1
	_, _
	bit3, mask3
)

func TestIota(t *testing.T) {

	// 说明 iota 表示 const 块的行索引，从0开始递增，每次加1
	// 也就是 同一行，iota代表的值是相同的
	fmt.Printf("bit0: %d mask0: %d\n", bit0, mask0)
	fmt.Printf("bit1: %d mask1: %d\n", bit1, mask1)
	fmt.Printf("bit3: %d mask3: %d\n", bit3, mask3)

	// 答案7，说明先计算左移再计算减法
	// 某些Go网站的优先级有误！！！
	// 例如 http://c.biancheng.net/view/5559.html
	// https://www.sojson.com/operation/go.html
	fmt.Println(1<<3 - 1)
	fmt.Println((1 << 3) - 1)
}
