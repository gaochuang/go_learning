/*
1.任意一个变量都包含pair<type, value>
2.它是有type指针和value指针组成
3.type包含静态的(static type)类型(例如:int、string、float)和具体(concrete type)类型(interface 所指的具体类型)
*/
package main

import (
	"fmt"
	"io"
	"os"
)

func stringPair() {
	var testStr string
	testStr = "abc"
	//pair<type:string, value:"abc">
	fmt.Println(testStr)

	//pair<type:, val: >
	var allType interface{}

	//将testStr赋值给allType
	//allType pair:<type: string, val: "abc">
	allType = testStr

	//allType断言, 相当于找pair中的type
	str, _ := allType.(string)
	fmt.Println(str)
}

// pair再连续赋值中会保持pair不变的
func openFileDemo() {
	//tty: pair<type: *os.File, value: "/dev/tty"文件描述符>，无论tty赋值给谁，pair不会改变
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		fmt.Println("open file failed")
		return
	}

	//r: pair<type: , value:>
	var r io.Reader
	//r: pari<type: *os.File, value: "/dev/tty"文件描述符>
	r = tty

	//w: pair<type:, val>
	var w io.Writer
	//将r强转成write,同时复制给w
	//: pair<type:*os.File, value: "/dev/tty"文件描述符>
	w = r.(io.Writer)

	w.Write([]byte("Hello this is test"))

}

type WriteBook interface {
	WriteBook()
}

type ReadBook interface {
	ReadBook()
}

type Book struct {
}

func (this *Book) WriteBook() {
	fmt.Println("write Book")
}

func (this *Book) ReadBook() {
	fmt.Println("read book")
}

func BookDemo() {

	//book: pair<type: Book, value: Book{}的地址>
	book := &Book{}

	//r此时type和value都是空: pair<type: , value: >
	var r ReadBook

	//赋值后r: pair<type: Book, value: Book{}的地址>
	r = book

	r.ReadBook()

	//w: pair<type:, value: >
	var w WriteBook

	w = r.(WriteBook) //此处可以断言成功是因为w和r的pair是一致的

	w.WriteBook()
}

func main() {
	stringPair()
	openFileDemo()
	BookDemo()
}
