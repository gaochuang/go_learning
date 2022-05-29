package ch3_test

import "testing"

type MyInt int64

//不支持隐性数据类型转换
func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	//b = a 将报错
	b = int64(a)

	var c MyInt //c 变量为自定义类型
	//c = b 报错 cannot use b (type int64) as type MyInt in assignment，别名是不同的类型，必须显示的转换
	c = MyInt(b) //显示转换
	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	var a = 1
	aPtr := &a

	//aPtr = aPtr+1错误
	*aPtr = 2
	t.Log(a)

	t.Logf("%T,%T", a, aPtr)
}

func TestString(t *testing.T) {
	var str string

	t.Log("*" + str + "*")
	t.Log(len(str))

}
