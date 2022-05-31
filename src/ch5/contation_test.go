package ch5_test

import (
	"fmt"
	"runtime"
	"testing"
)

func TestLoopTest(t *testing.T) {
	if a := 1 == 1; a {
		t.Log("1 == 1")
	}
	/*
		//someFunc多返回值
			if v, err := someFunc(); err{
				return
			}else{
				do something
			}
	*/
}

func TestSwitch1(t *testing.T) {

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("linux os.")
	case "windows":
		fmt.Println("windows os")
	default:
		fmt.Println(os)
	}

}

func TestSwitch2(t *testing.T) {
	for i := 0; i <= 10; i++ {
		switch v := 3; v {
		case 1, 3, 5:
			fmt.Println("1,3,5")
		case 2, 4, 6:
			fmt.Println("2,4,6")
		}
	}
}
