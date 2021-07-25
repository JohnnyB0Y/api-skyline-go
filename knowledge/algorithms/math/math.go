//  math.go
//
//
//  Created by JohnnyB0Y on 2021/07/26.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package math_test

import (
	"math/big"
)

/*
计数质数
统计所有小于非负整数 n 的质数的数量。
https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnzlu6/
*/

func countPrimes(n int) int {
	// 初始化
	zn := []int{2, 3, 5, 7}
	if n <= 11 {
		for i := len(zn) - 1; i >= 0; i-- {
			if n > zn[i] {
				return i + 1
			}
		}
		return 0
	}

	// 计算
	var i int64 = 11
	var num int64 = int64(n)
	for i <= num {
		if big.NewInt(i).ProbablyPrime(0) { // 判断是否素数
			zn = append(zn, int(i))
		}
		i += 2
	}

	N := len(zn)
	if zn[N-1] == n {
		return N - 1
	}
	return N
}
