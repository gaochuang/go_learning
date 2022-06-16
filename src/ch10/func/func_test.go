package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValue() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func TestFn(t *testing.T) {
	a, b := returnMultiValue()
	t.Log(a, b)
}

func SlowFn(n int) int {
	time.Sleep(5)
	return n
}

func functionSpentTime(runner func(n int) int) func(n int) int {

	return func(n int) int {
		startTime := time.Now()
		ret := runner(n)
		fmt.Printf("Spent time is %+v", time.Since(startTime).Seconds())
		return ret
	}
}

func TestFn2(t *testing.T) {
	fn := functionSpentTime(SlowFn)
	fmt.Println(fn(10))
}

func Sum(ops ...int) int {
	var ret int
	ret = 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
}

func TestDeferFunction(t *testing.T) {
	defer func() {
		t.Log("This is defer function")
	}()

	t.Log("Start")
	panic("panic err happend")

	t.Log("end") //不会被打印
}
