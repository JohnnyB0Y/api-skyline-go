//  math.go
//
//
//  Created by JohnnyB0Y on 2021/07/26.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package mymath

import (
	"math/big"

	"github.com/Workiva/go-datastructures/bitarray"
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

// 使用 bit array
func countPrimes4(n int) int {
	if n <= 2 {
		return 0
	}

	N := uint64(n)
	if N%2 != 0 { // 奇数变成偶数
		N += 1
	}

	// 去除 2 的倍数
	N = N / 2
	isNotPrimes := bitarray.NewBitArray(uint64(N))
	isNotPrimes.SetBit(0) // 下标0为1，是非质数

	count := 1 // 这个是把2计算在内
	var p uint64 = 1
	for ; p*p < N; p++ {
		// 把非素数标出来
		notPrime, _ := isNotPrimes.GetBit(p)
		if !notPrime {
			count++ // 素数加起来
			step := p*2 + 1
			for i := p + step; i < N; i += step {
				// 1 -> 4 -> 7 -> 10
				// 2 -> 7 -> 12 ->
				isNotPrimes.SetBit(i)
			}
		}
	}

	// 统计后半段的质数
	var lastPrime uint64 = 1 // 这个是 3 的下标
	for ; p < N; p++ {
		notPrime, _ := isNotPrimes.GetBit(p)
		if !notPrime {
			count++ // 素数加起来
			lastPrime = p
		}
	}

	// 判断 n是否质数，按题目排除
	if lastPrime*2+1 == uint64(n) {
		count-- // 去除自己
	}

	return count
}

/**
3的幂
给定一个整数，写一个函数来判断它是否是 3 的幂次方。如果是，返回 true ；否则，返回 false 。

整数 n 是 3 的幂次方需满足：存在整数 x 使得 n == 3x

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnsdi2/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func isPowerOfThree(n int) bool {

	// 1 3 9 27 81 243 729 2187 6561 19683 59049 177147 531441
	nums3 := map[int]bool{
		1: true, 729: true,
		3: true, 2187: true,
		9: true, 6561: true,
		27: true, 19683: true,
		81: true, 59049: true,
		243: true, 177147: true,
		531441: true, 1594323: true,
		4782969: true, 14348907: true,
		43046721: true, 129140163: true,
		387420489: true, 1162261467: true,
	}

	for i := 1162261467; i <= n; i *= 3 {
		nums3[i] = true
	}

	return nums3[n]
}

func isPowerOfThree2(n int) bool {
	if n < 1 {
		return true
	}
	// 参考答案！！！
	return 1162261467%n == 0
}
