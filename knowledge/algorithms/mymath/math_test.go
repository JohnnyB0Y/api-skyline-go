//  math_test.go
//
//
//  Created by JohnnyB0Y on 2021/07/26.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package mymath

import "testing"

func TestCountPrimes(t *testing.T) {
	t.Log(countPrimes(100))
	t.Log(countPrimes2(100))
	t.Log(countPrimes3(100))
	t.Log(countPrimes4(100))

	t.Log(countPrimes(7))
	t.Log(countPrimes2(7))
	t.Log(countPrimes3(7))
	t.Log(countPrimes4(7))
}

func TestIsPowerOfThree(t *testing.T) {
	t.Log(isPowerOfThree(135005697))
}

func BenchmarkCountPrimes3(b *testing.B) {
	// 100内：BenchmarkCountPrimes3-8   	10797572	      99.8 ns/op	      64 B/op	       1 allocs/op
	// 100万：BenchmarkCountPrimes3-8   	     250	   4786838 ns/op	  507905 B/op	       1 allocs/op
	for i := 0; i < b.N; i++ {
		countPrimes3(100)
	}
}

func BenchmarkCountPrimes4(b *testing.B) {
	// 100内：BenchmarkCountPrimes4-8   	 2642569	       455 ns/op	      56 B/op	       2 allocs/op
	// 100万：BenchmarkCountPrimes4-8   	     150	   7077849 ns/op	   65584 B/op	       2 allocs/op
	for i := 0; i < b.N; i++ {
		countPrimes4(100)
	}
}

/**
countPrimes3 用的是bool数组，性能比较优秀；
countPrimes4 用的是 bit array，内存比较优秀；


bit array 性能比 bool数组 低不少，但内存占用上比 bool数组 优秀，理论上内存占用只有 bool数组的 1/8;
	- 由于主存访问以字节为基本单位，处理bit位增加了不少CPU消耗；
	- 如果程序对内存比较敏感，那么使用 bit array 是不错的选择；
	- 不然的话，没有必要优化到 bit 位；
*/
