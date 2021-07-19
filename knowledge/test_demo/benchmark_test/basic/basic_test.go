//  basic_test.go
//
//
//  Created by JohnnyB0Y on 2021/07/19.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package basic

import (
	"fmt"
	"testing"
)

var gs string

// BenchmarkSprint tests the performance of using Sprint.
func BenchmarkSprint(b *testing.B) {
	var s string

	for i := 0; i < b.N; i++ {
		s = fmt.Sprint("hello")
	}

	gs = s
}

// BenchmarkSprint tests the performance of using Sprintf.
func BenchmarkSprintf(b *testing.B) {
	var s string

	for i := 0; i < b.N; i++ {
		s = fmt.Sprintf("hello")
	}

	gs = s
}
