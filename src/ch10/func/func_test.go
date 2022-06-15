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

func slowFunc(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

//类似装饰者模式
func FuntionSpend(runner func(op int) int) func(op int) int {
	return func(n int) int {
		startTime := time.Now()
		ret := runner(n)
		fmt.Println("time spend: ", time.Since(startTime).Seconds())
		return ret
	}
}

//函数式编程
func TestFn1(t *testing.T) {
	tsSF := FuntionSpend(slowFunc)
	t.Log(tsSF(10))
}
