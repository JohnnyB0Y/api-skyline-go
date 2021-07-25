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
寻找 100 以内质数时，countPrimes3 比较优秀；
当 寻找的质数提高到 100万 时，countPrimes4 性能和内存比较优秀；

countPrimes3 用的是bool数组，countPrimes4 用的是 bit array；
*/
