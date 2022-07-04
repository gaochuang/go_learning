/*
go routine并发模式 - 仅需任意一个任务完成

场景：使用百度浏览器和谷歌浏览器同时搜索一个关键字，任意
一个成功即可
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
	return fmt.Sprintf("The result is from %d", i)
}

func FirstResponse() string {

	numberRunner := 10
	//使用无缓冲的channel,协程向channel写数据，如果channel没有接收者，会导致阻塞
	//ch := make(chan string)

	//有缓冲的buffer,不必要等待接收者接收就可以写入
	ch := make(chan string, numberRunner)

	for i := 0; i < numberRunner; i++ {
		go func(input int) {
			ret := runTask(input)
			ch <- ret
		}(i)

	}
	return <-ch

}

func TestFirstResponse(t *testing.T) {

	t.Log("Beforr", runtime.NumGoroutine())
	t.Log(FirstResponse())
	time.Sleep(time.Second * 1)
	t.Log("After", runtime.NumGoroutine())
}
