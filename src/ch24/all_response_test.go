/*
go routine并发模式 - 等待所有的任务完成
*/

package ch24_test

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(i int) string {
	time.Sleep(time.Millisecond * 10)
	return fmt.Sprintf("this is task %d finish", i)
}

func allResponse() string {

	numRunner := 10
	ch := make(chan string, numRunner)
	for i := 0; i < numRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)

	}

	finalResult := ""
	for i := 0; i < numRunner; i++ {
		finalResult += <-ch + "\n"
	}

	return finalResult

}

func TestAllResponse(t *testing.T) {
	t.Log("Before", runtime.NumGoroutine())
	t.Log(allResponse())
	t.Log("After", runtime.NumGoroutine())
}
