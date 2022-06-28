package ch19_test

import (
	"testing"
	"time"
)

func TestSelect1(t *testing.T) {

	ch1 := make(chan int)
	ch2 := make(chan int)

	select {
	case <-ch1:
		t.Log("trigger ch1")
	case <-ch2:
		t.Log("trgeer ch2")
	default:
		t.Log("error")
	}
}

func TestTimeout(t *testing.T) {
	ch1 := make(chan int)
	select {
	case <-ch1:
		t.Log("ch1")
	case <-time.After(time.Millisecond * 5):
		t.Log("time out")

	}
}
