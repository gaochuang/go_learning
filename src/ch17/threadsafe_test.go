package ch17_test

import (
	"sync"
	"testing"
)

func TestThreadSafe(t *testing.T) {

	counter := 0
	var mux sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 5000; i++ {
		wg.Add(1) //增加要等待的量
		go func() {
			defer func() {
				mux.Unlock()
				wg.Done() //等待的任务结束，告诉waitGroup
			}()
			//加锁原因是因为counter在不同的并发goroutine是共享的，不是线程安全的
			mux.Lock()
			counter++
		}()
	}

	wg.Wait() //wait到所有要等待的任务完成
	t.Log("counter: ", counter)
}
