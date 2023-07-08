package main

import "fmt"

type Student struct {
	//如果类的属性首字母大写，表示该属性对外可以访问，否则的话只能够结构体的内部访问
	Name  string
	Age   int
	Score int32
}

// 方法绑定结构体
func (this *Student) GetName() (name string) {
	fmt.Println("Name = ", this.Name)
	name = this.Name
	return
}

func (this *Student) SetName(name string) {
	this.Name = name
}

//=================继承==================

type Human struct {
	Name string
	Sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human eat......")
}

func (this *Human) Walk() {
	fmt.Println("Human walk.......")
}

// SuperMan 继承Human
type SuperMan struct {
	Human
	level int
}

// 子类，superman的新方法
func (this *SuperMan) Fly() {
	fmt.Println("I'm sumpperMan, I can fly")
}

//子类 superman 重写父类Walk方法
func (this *SuperMan) Walk() {
	fmt.Println("I'm supperman, I can fly, needn't walk")
}

func main() {
	var stu1 Student
	stu1.Age = 11
	stu1.Name = "Bob"
	stu1.Score = 100

	fmt.Println("stu1: ", stu1)

	stu2 := Student{Name: "David", Age: 12, Score: 9999}
	fmt.Println("stu2: ", stu2)

	man := Human{Name: "BoB", Sex: "male"}
	man.Walk()

	var super SuperMan
	super.Name = " man 1"
	super.level = 0

	super.Fly()
	super.Walk()

}
