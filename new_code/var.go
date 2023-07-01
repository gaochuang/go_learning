package main

import "fmt"

//声明全局变量： 方法一、二、三可以

var g_A int
var g_B int = 12
var g_C = "global val"

//error
// g_D := 10 不能使用，只能在函数体内

//四种变量声明
func main() {

	//方法一: 声明一个变量默认值是0
	var a int
	fmt.Printf("a type %T, value: %+v\n", a, a)

	//方法二: 声明一个变量，给一个默认值
	var b string = "hello golang"
	fmt.Printf("b type %T, value: %+v\n", b, b)

	//方法三：在初始化时候，可以省去数据类型，通过值自动匹配当前变量的数据类型,不建议使用
	var c = 1000
	fmt.Printf("c type %T, value: %+v\n", c, c)

	//方法四： 省去var关键字，直接自动匹配，常用

	d := "abc"
	e := 123
	f := 0.123

	type Info struct {
		name string
	}

	fmt.Printf("d type %T, value: %+v\n", d, d)
	fmt.Printf("e type %T, value: %+v\n", e, e)
	fmt.Printf("f type %T, value: %+v\n", f, f)

	fmt.Printf("Info type %T, value: %+v\n", Info{}, Info{})

	fmt.Printf("g_A type %T, value: %+v\n", g_A, g_A)
	fmt.Printf("g_B type %T, value: %+v\n", g_B, g_B)
	fmt.Printf("g_C type %T, value: %+v\n", g_C, g_C)

	//声明多个变量
	var aa, bb int = 100, 200

	fmt.Println(aa, bb)

	var cc, dd = 100, "test"
	fmt.Println(cc, dd)

	//多行多变量声明
	var (
		kk int  = 100
		ff bool = true
	)

	fmt.Println(kk, ff)

}
