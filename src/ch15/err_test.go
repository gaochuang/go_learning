package ch15_test

import (
	"errors"
	"testing"
)

var LessThanTwoError = errors.New("n should be not less than 2")
var LagrgerThenHundredError = errors.New("n should be not lager than 100")

func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, LagrgerThenHundredError
	}

	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	if v, err := GetFibonacci(5); err != nil {
		if err == LessThanTwoError {
			t.Log(LessThanTwoError)
		} else {
			t.Log(v)
		}
	}

}
