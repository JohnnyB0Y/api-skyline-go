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
