package lib1

//Lib1包中导入Lib2,先打印 init lib2

import "import_test/lib2"
import "fmt"

func Lib1Test() {
	lib2.Lib2Test()
	fmt.Println("Lib1Test()...")
}
func init() {
	fmt.Println("init lib1")
}
