package unittest

import "testing"

func TestAdd(t *testing.T) {
	var a, b int = 1, 2
	var expected int = 3

	var actual int = Add(a, b)
	if actual != expected {
		t.Errorf("Add(%d, %d) = %d; expected: %d", a, b, actual, expected)
	}
}

/**
执行测试
go test


执行测试，打印测试详细信息
go test -v


执行测试，打印测试详细信息 和 测试覆盖率
go test -v -cover

*/
