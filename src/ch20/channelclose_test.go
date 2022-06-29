package ch20_test

import (
	"fmt"
	"sync"
	"testing"
)

/*
func TestWriteToCloseChannel(t *testing.T) {
	ch1 := make(chan int)
	close(ch1)

	ch1 <- 1
}
*/

func provider(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()

	}()

}

func consumer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			if v, ok := <-ch; ok {
				fmt.Printf("%+v\n", ok)
				fmt.Println("recv", v)
			} else {
				fmt.Printf("%+v\n", ok)
				break
			}
		}
		wg.Done()
	}()

}

func TestCP(t *testing.T) {
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	provider(ch, &wg)
	wg.Add(1)
	consumer(ch, &wg)
	wg.Wait()
}
