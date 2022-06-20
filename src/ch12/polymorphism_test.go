package ch12_test

import (
	"fmt"
	"testing"
)

tye Code string
type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {
}

func (p *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(hello world)"
}

type JavaProgrammer struct {
}

func (p *JavaProgrammer) WriteHelloWorld() Code {
	return "java.out.print(hello world)"
}
/*入参必须是指针*/
func WriteFirstProgrammer(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func TestPolymorphism(t *testing.T) {
	goprammer := new(GoProgrammer)
	javaprommer := new(JavaProgrammer)
	/*传入的必须是指针*/
	WriteFirstProgrammer(goprammer)
	WriteFirstProgrammer(javaprommer)
}
