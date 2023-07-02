//函数多返回值

package main

import "fmt"

func foo1(a string, b int) int {
	fmt.Println(a)
	fmt.Println(b)
	c := 100
	return c
}

//返回多个返回值，都是匿名
func foo2(a int, b string) (int, int) {
	fmt.Println(a)
	fmt.Println(b)

	return 666, 777
}

//返回多个返回值，有名称
func foo3(a int) (ret1 int, ret2 int) {
	fmt.Println(a)

	//ret1, ret2属于形参，作用域空间是函数体内部空间
	fmt.Printf("%T\n", ret1)
	fmt.Printf("%T\n", ret2)

	//初始化值是0
	fmt.Println(ret1)
	fmt.Println(ret2)

	//给有名称的返回值赋值
	ret1 = 1000
	ret2 = 2000
	return
}

//两个返回值相同，可以合二为一
func foo4() (ret1, ret2 int) {
	ret1 = 100
	ret2 = 200
	return
}

func main() {
	fmt.Println(foo1("a", 10))
	fmt.Println(foo2(1, "dd"))
	fmt.Println(foo3(3))

	fmt.Println(foo4())
}
