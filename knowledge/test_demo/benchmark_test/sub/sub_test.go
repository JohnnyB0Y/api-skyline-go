//  sub_test.go
//
//
//  Created by JohnnyB0Y on 2021/07/15.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

// go test -run none -bench . -benchtime 3s -benchmem
// go test -run none -bench BenchmarkSprint/none -benchtime 3s -benchmem
// go test -run none -bench BenchmarkSprint/format -benchtime 3s -benchmem

// Basic sub-benchmark test.
// 子基准测试
package basic

import (
	"fmt"
	"testing"
)

var gs string

// BenchmarkSprint tests all the Sprint related benchmarks as
// sub benchmarks.
func BenchmarkSprint(b *testing.B) {
	b.Run("none", benchSprint)
	b.Run("format", benchSprintf)
}

// benchSprint tests the performance of using Sprint.
func benchSprint(b *testing.B) {
	var s string

	for i := 0; i < b.N; i++ {
		s = fmt.Sprint("hello")
	}

	gs = s
}

// benchSprintf tests the performance of using Sprintf.
func benchSprintf(b *testing.B) {
	var s string

	for i := 0; i < b.N; i++ {
		s = fmt.Sprintf("hello")
	}

	gs = s
}
