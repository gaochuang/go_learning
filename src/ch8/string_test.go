package ch8_test

import (
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	var s string //s是一个数据类型不是一个地址或者引用
	t.Log(s)     //空
	s = "hello"
	t.Log(s, len(s))

	s = "\xE4\xB8\xA5" //	可以存储二进制数据
	t.Log(s)           //day

	s = "中"
	t.Log(len(s))  //3 因为一个unicode占3个byte
	c := []rune(s) //1 转成1个unicode
	t.Logf("rune %d", len(c))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)
}
func TestUnicode(t *testing.T) {
	s := "这个是测试"
	for _, c := range s {
		t.Logf("%[1]c, %[1]x", c)
	}
}
func TestStringfn(t *testing.T) {
	var s string = "A,B,C"
	parts := strings.Split(s, ",") //字符串分割，以","进行分割
	for index, v := range parts {
		t.Logf("index: %d, value: %s", index, v)
	}
	s1 := strings.Join(parts, "-")
	t.Log(s1) //A-B-C

	s3 := []string{"D", "P", "M"}

	s4 := strings.Join(s3, "?") //D?P?M
	t.Log(s4)

}
