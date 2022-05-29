package fib

import "testing"

func TestFib(t *testing.T) {
	var (
		a int = 1
		b int = 1
	)
	t.Log(a)
	for i := 1; i < 5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + b
	}
}

func TestValue(t *testing.T) {
	const (
		Monday = iota + 1 //1. iota在const关键字出现时将被重置为0; 2.const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)
		Tuesday
		Wednesday
		Thursday
		Frinday
		Saturday
		Sunday
	)

	t.Log(Monday) //1
	t.Log(Sunday) //7
}

func TestBitValue(t *testing.T) {
	const (
		ReadOnly = 1 << iota
		WriteOnly
		ExeCOnly
	)
	t.Log(ReadOnly)
	t.Log(WriteOnly)
	t.Log(ExeCOnly)
}
