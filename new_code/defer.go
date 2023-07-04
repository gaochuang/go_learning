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

/*
recover错误拦截:
1. Go语言提供拦截运行时pannic的内建函数recover。它可以是当前的程序从pannic的状态中恢复并重新获得流程控制权
#### recover函数只有在defer函数中调用才有效#####
*/

func pannicAndRecover() {

	//错误拦截要在产生错误前设置
	defer func() {
		//设置recover拦截错误信息
		err := recover()
		if err != nil {
			//产生pannic异常，打印错误信息
			fmt.Println("### will print pannic information ####")
			fmt.Println(err)
		}
	}()

	//主动产生一个pannic
	panic("test pannic")
}

func main() {
	defer func3()
	defer func2()
	defer func1()

	fmt.Println(returnAndDefer())

	pannicAndRecover()
}
