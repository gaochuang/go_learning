package ch13_test

import (
	"fmt"
	"testing"
)

type EmptyInterface interface {
}

func getType1(input EmptyInterface) {
	if v, ok := input.(int); ok {
		fmt.Printf("integer %d\n", v)
		return
	}

	if v, ok := input.(string); ok {
		fmt.Printf("string %s\n", v)
		return
	}

	fmt.Println("unknow")
}

func getType2(input EmptyInterface) {
	switch p := input.(type) { //断言
	case int:
		fmt.Printf("integer %d\n", p)
	case string:
		fmt.Printf("string %s\n", p)
	default:
		fmt.Println("Unknown")
	}
}

func TestEmpryInterface(t *testing.T) {
	getType1(222)
	getType1("testy")
	getType1(1.0)

	getType2(222)
	getType2("testy")
	getType2(1.0)
}
