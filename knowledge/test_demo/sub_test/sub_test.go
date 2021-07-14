package subtest

import (
	"testing"
	"time"
)

func subtest1(t *testing.T) {
	var a = 1
	var b = 1
	var expected = 2
	actual := Add(a, b)
	if actual != expected {
		t.Errorf("Add(%d, %d = %d; expected: %d", a, b, actual, expected)
	}
}

func subtest2(t *testing.T) {
	var a = 1
	var b = 2
	var expected = 3
	actual := Add(a, b)
	if actual != expected {
		t.Errorf("Add(%d, %d = %d; expected: %d", a, b, actual, expected)
	}
}

func subtest3(t *testing.T) {
	var a = 1
	var b = 3
	var expected = 4
	actual := Add(a, b)
	if actual != expected {
		t.Errorf("Add(%d, %d = %d; expected: %d", a, b, actual, expected)
	}
}

func parallelTest1(t *testing.T) {
	t.Parallel()
	time.Sleep(3 * time.Second)
	// do something
}

func parallelTest2(t *testing.T) {
	t.Parallel()
	time.Sleep(2 * time.Second)
	// do something
}

func parallelTest3(t *testing.T) {
	t.Parallel()
	time.Sleep(1 * time.Second)
	// do something
}

func TestSub(t *testing.T) {
	// setup
	t.Log("Test Sub Setup...")

	// 子测试名称  子测试函数
	t.Run("A=1", subtest1)
	t.Run("A=2", subtest2)
	t.Run("B=1", subtest3)

	// finished
	t.Log("Test Sub Finished...")
}

func TestSubParallel(t *testing.T) {
	// setup
	t.Log("Test Parallel Setup...")

	// 并发测试
	t.Run("group", func(t *testing.T) {
		t.Run("test1", parallelTest1)
		t.Run("test2", parallelTest2)
		t.Run("test3", parallelTest3)
	})

	// finished
	t.Log("Test Parallel Finished...")
}

/**
执行当前所有测试
go test -v -cover

选择测试过滤语法
// 选择A=1 的子测试 执行
go test sub_test.go sub.go -v -run Sub/A=1

// 选择所有A= 的子测试 执行
go test sub_test.go sub.go -v -run Sub/A=

// 选择所有B= 的子测试 执行
go test sub_test.go sub.go -v -run Sub/B=

*/
