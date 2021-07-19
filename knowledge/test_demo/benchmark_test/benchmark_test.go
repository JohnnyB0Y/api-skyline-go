//  benchmark_test.go
//
//
//  Created by JohnnyB0Y on 2021/07/19.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package benchmarktest

import "testing"

func BenchmarkMakeSliceWithoutPreallocatedMemory(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MakeSliceWithoutPreallocatedMemory(100000)
	}
}

func BenchmarkMakeSliceWithPreallocatedMemory(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MakeSliceWithPreallocatedMemory(100000)
	}
}

/**
执行基准测试
go test -bench=.



*/
