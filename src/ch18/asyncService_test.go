package ch18_test

import (
	"fmt"
	"testing"
	"time"
)

func Service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func OtherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("working done")

}

func AsyncService() chan string {
	retCh := make(chan string)
	go func() {
		ret := Service()
		fmt.Println("returned result")
		retCh <- ret
		fmt.Println("Service exited")
	}()
	return retCh
}

func TestAsyncService(t *testing.T) {
	retch := AsyncService()
	OtherTask()
	fmt.Println(<-retch)
}
