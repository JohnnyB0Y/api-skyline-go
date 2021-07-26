//  math.go
//
//
//  Created by JohnnyB0Y on 2021/07/26.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package mymath

import (
	"math/big"
	"strconv"

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

func isPowerOfThree3(n int) bool {
	for i := 1; i <= n; i *= 3 {
		if i == n {
			return true
		}
	}
	return false
}

/**
Fizz Buzz
写一个程序，输出从 1 到 n 数字的字符串表示。

1. 如果 n 是3的倍数，输出“Fizz”；

2. 如果 n 是5的倍数，输出“Buzz”；

3.如果 n 同时是3和5的倍数，输出 “FizzBuzz”。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xngt85/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func fizzBuzz(n int) []string {
	str := make([]string, n)
	for i := 1; i <= n; i++ {

		switch {
		case i%3 == 0 && i%5 == 0:
			str[i-1] = "FizzBuzz"
		case i%3 == 0:
			str[i-1] = "Fizz"
		case i%5 == 0:
			str[i-1] = "Buzz"
		default:
			str[i-1] = strconv.Itoa(i)
		}
	}
	return str
}

/**
罗马数字转整数
罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。

字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
给定一个罗马数字，将其转换成整数。输入确保在 1 到 3999 的范围内。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn4n7c/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

const (
	I    = 'I'
	V    = 'V'
	X    = 'X'
	L    = 'L'
	C    = 'C'
	D    = 'D'
	M    = 'M'
	STOP = 'S'
)

type romanHelper struct {
	sum    int  // 总和
	subSum int  // 累计
	ps     rune // 当前累计的符号
}

func (h *romanHelper) subSumAdd(num int, ps rune) {
	h.subSum += num
	h.ps = ps
}

func (h *romanHelper) sumDelelte(num int) {
	h.sum += h.subSum - num
	h.subSum = 0
	h.ps = STOP
}

func (h *romanHelper) sumAdd(num int) {
	h.sum += h.subSum + num
	h.subSum = 0
	h.ps = STOP
}

func romanToInt(s string) int {

	helper := &romanHelper{0, 0, STOP}
	// 从后往前扫
	for i := len(s) - 1; i >= 0; i-- {
		switch s[i] {
		case I:
			if helper.ps == STOP || helper.ps == I {
				helper.subSumAdd(1, I) // 累计
			} else if helper.ps == V || helper.ps == X { // IV XV
				helper.sumDelelte(1)
			}
		case V:
			if helper.ps == STOP {
				helper.subSumAdd(5, V) // 累计
			} else if helper.ps == I { // VI
				helper.sumAdd(5)
			}
		case X:
			if helper.ps == STOP || helper.ps == X || helper.ps == V {
				helper.subSumAdd(10, X) // 累计
			} else if helper.ps == I { // XI
				helper.sumAdd(10)
			} else if helper.ps == L || helper.ps == C { // XL XC
				helper.sumDelelte(10)
			}
		case L:
			if helper.ps == STOP {
				helper.subSumAdd(50, L) // 累计
			} else if helper.ps == X { // LX
				helper.sumAdd(50)
			}
		case C:
			if helper.ps == STOP || helper.ps == C || helper.ps == L || helper.ps == V {
				helper.subSumAdd(100, C) // 累计
			} else if helper.ps == X { // CX
				helper.sumAdd(100)
			} else if helper.ps == D || helper.ps == M { // CD CM
				helper.sumDelelte(100)
			}
		case D:
			if helper.ps == STOP {
				helper.subSumAdd(500, D) // 累计
			} else if helper.ps == C { // DC
				helper.sumAdd(500)
			}
		case M:
			if helper.ps == STOP || helper.ps == M || helper.ps == D || helper.ps == L || helper.ps == V {
				helper.subSumAdd(1000, M) // 累计
			} else if helper.ps == C { // MC
				helper.sumAdd(1000)
			}
		}
	}
	helper.sumAdd(0)
	return helper.sum
}
