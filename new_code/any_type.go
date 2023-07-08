package main

import "fmt"

//interface{}是万能类型
func valType(arg interface{}) {
	fmt.Println(arg)

	//interface{}断言机制
	val, ok := arg.(string)
	if ok {
		fmt.Printf("arg type %T\n", val)
		fmt.Println("arg value is ", val)
	} else {
		fmt.Println("arg is not string")
	}

	switch arg.(type) {
	case nil:
		fmt.Println("arg is nil")
	case string:
		fmt.Println("arg is string")
	case int32:
		fmt.Println("arg is int32")
	case bool:
		fmt.Println("arg is bool")
	default:
		fmt.Println("arg is unknown")

	}
}

func main() {
	valType("123")
	valType(1)
	valType(true)
	valType([]int{1, 2, 3})
}
