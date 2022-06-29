package ch21_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func isCancel(ch chan struct{}) bool {
	select {
	case <-ch:
		return true
	default:
		return false
	}
}

func cancel(ch chan struct{}) {
	close(ch)
}

func TestCancel(t *testing.T) {

	var wg sync.WaitGroup
	ch := make(chan struct{})

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(i int, ch chan struct{}) {
			defer func() {
				wg.Done()
			}()
			for {
				if isCancel(ch) {
					break
				} else {
					time.Sleep(time.Millisecond * 5)
				}
			}
			fmt.Printf("goroutine %d  is cancel \n", i)

		}(i, ch)
	}

	time.Sleep(time.Millisecond * 100)

	cancel(ch)
	fmt.Println("END")
	wg.Wait()
	fmt.Println("END")
}
