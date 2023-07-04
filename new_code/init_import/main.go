package main

//main包中再次导入lib2包，不会打印init lib2，因为相同的包import多次，实际只导入一次
import (
	"import_test/lib1"
	"import_test/lib2"
	_ "import_test/lib2" //匿名导包，可以只使用lib2的init方法，不适用其它方法
	//  . "import_test/lib2" 把整个包导入到当前包
)

func main() {
	lib1.Lib1Test()
	lib2.Lib2Test()
}
