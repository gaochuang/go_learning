package ch12_test

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (f *Pet) Speak() {
	fmt.Println("Seak")
}

func (f *Pet) Eat() {
	fmt.Println("Eat")
}

type Dog struct {
	Pet //匿名嵌套类型
}

func (d *Dog) Name() {
	fmt.Println("Name")
}

func TestExtension(t *testing.T) {

	dog := new(Dog)
	dog.Speak()
	dog.Eat()
	dog.Name()
}
