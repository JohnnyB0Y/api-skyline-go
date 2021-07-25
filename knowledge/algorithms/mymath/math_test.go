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

	t.Log(countPrimes(7))
	t.Log(countPrimes2(7))
	t.Log(countPrimes3(7))
}
