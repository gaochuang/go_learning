package ch16_test

import (
	"errors"
	"testing"
)

func TestPanic(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log(err)
		}
	}()

	t.Log("panic test")
	panic(errors.New("panic happend"))
}
