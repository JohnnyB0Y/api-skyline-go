package memory_test

import (
	"testing"
	"unsafe"
)

func TestMakeNew(t *testing.T) {
	m := make([]int, 2)
	n := new([]int)
	s := []int{}
	a := [...]int{1, 2, 3}
	a2 := [...]int{1, 2, 3}
	i := 10086
	i2 := 10086
	i3 := 113114
	str := "str"
	str2 := "str"
	str3 := "bbq"
	m2 := make([]int, 2)
	n2 := new([]int)

	t.Logf("make pointer: %p\n", &m)
	t.Logf("new pointer: %p\n", n)
	t.Logf("slice pointer: %p\n", &s)
	t.Logf("array 1 pointer: %p\n", &a)
	t.Logf("array 2 pointer: %p\n", &a2)

	t.Logf("int 1 pointer: %p\n", &i)
	t.Logf("int 2 pointer: %p\n", &i2)
	t.Logf("int 3 pointer: %p\n", &i3)

	t.Logf("str 1 pointer: %p\n", &str)
	t.Logf("str 2 pointer: %p\n", &str2)
	t.Logf("str 3 pointer: %p\n", &str3)

	t.Logf("make 2 pointer: %p\n", &m2)
	t.Logf("new 2 pointer: %p\n", n2)

	// 算一下地址，看是在栈上还是堆上
	pm := uintptr(unsafe.Pointer(&m))
	pn := uintptr(unsafe.Pointer(n))
	ps := uintptr(unsafe.Pointer(&s))
	pa := uintptr(unsafe.Pointer(&a))
	pi := uintptr(unsafe.Pointer(&i))
	pm2 := uintptr(unsafe.Pointer(&m2))
	t.Logf("pi - ps = %d; ps - pn = %d; pn - pm = %d;\n", pi-ps, ps-pn, pn-pm)

	t.Logf("pm2 - pm = %d\n", pm2-pm)
	t.Logf("pi - pa = %d\n", pi-pa)
}
