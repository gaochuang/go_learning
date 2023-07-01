package main

import "fmt"

//const 定义枚举类型
const (
	// 在const中使用iota，每行的iota都会增加1，第一行iota默认是0
	BEJING    = iota // 0
	SHNANGHAI        //iota = 1
	SHENZHEN         //iota = 2
)

const (
	// 在const中使用iota，每行的iota都会增加1，第一行iota默认是0
	JINAN     = 10 * iota // 0
	QINGDAO               //10 = 10 * iota(1)
	ZAOZHUANG             //20 = 10 * iota(2)
)

//第一行给定公式，其它行将继承公式，除非改变公式
const (
	a, b = iota + 1, iota + 2 // iota = 0, a = iota + 1, b = iota + 2， a = 1,  b= 2
	c, d                      //iota = 1, c = iota + 2, d = iota + 2, c = 2, d = 3
	e, f                      //iota = 2, e = iota + 1, f = iota + 2, e = 3, 4
	g, h = iota * 2, iota * 3 //iota = 3, g = iota*2, h = iota*3, g = 6, h = 9
	i, k                      //iota = 4, i = iota*2 , k = iota*3, i = 8, k = 12
)

func main() {

	//常量(只读属性)
	const length int = 10

	//length = 2 不能给常量负责，不允许修改

	fmt.Println(BEJING, SHNANGHAI, SHENZHEN)
	fmt.Println(JINAN, QINGDAO, ZAOZHUANG)

	fmt.Println(a, b, c, d, e, f, g, h, i, k)

	//iota只能配合const()使用
	//var w int = iota
}
