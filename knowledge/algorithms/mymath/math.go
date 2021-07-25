//  math.go
//
//
//  Created by JohnnyB0Y on 2021/07/26.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package mymath

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

func countPrimes2(n int) int {

	zn := SieveOfEratosthenes(n)
	N := len(zn)

	if N < 1 {
		return 0
	}

	if zn[N-1] == n {
		return N - 1
	}
	return N
}

// https://siongui.github.io/2017/04/17/go-sieve-of-eratosthenes/
func SieveOfEratosthenes(num int) []int {

	if num < 2 {
		return []int{}
	}

	if num%2 != 0 { // 奇数变成偶数
		num += 1
	}

	// 去除 2 的倍数
	N := num / 2
	isNotPrimes := make([]bool, N)
	isNotPrimes[0] = true

	//[0  1  2  3  4  5  6   7  8  9   10  11  12  13 14  15 16   17   18  19  20] 下标
	//+1 +2 +3 +4 +5 +6  +7 +8 +9 +10
	// 1  3  5  7  9  11 13 15 17  19  21  23  25  27 29  31 33
	//    i  j  k  i        ij         ik      j

	for p := 1; p*p < N; p++ {
		// 把非素数标出来
		if !isNotPrimes[p] {
			step := p*2 + 1
			for i := p + step; i < N; i += step {
				// 1 -> 4 -> 7 -> 10
				// 2 -> 7 -> 12 ->
				isNotPrimes[i] = true
			}
		}
	}

	// 组装素数
	var primes = []int{2}
	for p := 1; p < N; p++ {
		if !isNotPrimes[p] {
			primes = append(primes, p*2+1)
		}
	}
	return primes
}

func countPrimes3(n int) int {
	if n <= 2 {
		return 0
	}

	N := n
	if N%2 != 0 { // 奇数变成偶数
		N += 1
	}

	// 去除 2 的倍数
	N = N / 2
	isNotPrimes := make([]bool, N)
	isNotPrimes[0] = true

	count := 1 // 这个是把2计算在内
	p := 1
	for ; p*p < N; p++ {
		// 把非素数标出来
		if !isNotPrimes[p] {
			count++ // 素数加起来
			step := p*2 + 1
			for i := p + step; i < N; i += step {
				// 1 -> 4 -> 7 -> 10
				// 2 -> 7 -> 12 ->
				isNotPrimes[i] = true
			}
		}
	}

	// 统计后半段的质数
	lastPrime := 1 // 这个是 3 的下标
	for ; p < N; p++ {
		if !isNotPrimes[p] {
			count++ // 素数加起来
			lastPrime = p
		}
	}

	// 判断 n是否质数，按题目排除
	if lastPrime*2+1 == n {
		count-- // 去除自己
	}

	return count
}
