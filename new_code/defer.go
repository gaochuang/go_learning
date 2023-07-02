package main

import "fmt"

func func1() {
	fmt.Println("func1()...")
}

func func2() {
	fmt.Println("func2()...")
}

func func3() {
	fmt.Println("func3()...")
}

/*
defer执行顺序 压栈和出栈方式
*/

/*
defer和return执行顺序
*/

func deferFunc() int {
	fmt.Println("deferFunc called...")
	return 0
}

func returnFunc() int {
	fmt.Println("return  func called...")
	return 0
}

//returnFunc函数早于deferFunc函数执行
func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}

func main() {
	defer func3()
	defer func2()
	defer func1()

	fmt.Println(returnAndDefer())
}
