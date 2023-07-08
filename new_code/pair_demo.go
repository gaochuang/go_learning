/*
1.任意一个变量都包含pair<type, value>
2.它是有type指针和value指针组成
3.type包含静态的(static type)类型(例如:int、string、float)和具体(concrete type)类型(interface 所指的具体类型)
*/
package main

import "fmt"

func stringPair() {
	var testStr string
	testStr = "abc"
	//pair<type:string, value:"abc">
	fmt.Println(testStr)
}

func main() {
	stringPair()
}
