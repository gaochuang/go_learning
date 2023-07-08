/*
多态: 同一个行为具有多个不同的表现形式。也就是一个类的实例(对象)的相同方法在不同场景下有不同的表现形式。

*/

package main

import "fmt"

//interface 本质是一个指针
type Humans interface {
	GetJob()
	GetName() string
	GetSkill()
}

//Teacher继承Human接口，只需要实现父类的方法
type Teacher struct {
	Name  string
	Skill string
}

func (this *Teacher) GetName() string {
	return this.Name
}
func (this *Teacher) GetJob() {
	fmt.Println("I'm a teacher")
}

func (this *Teacher) GetSkill() {
	fmt.Println("I'm a teacher, my skill is fly")
}

type Students struct {
	Name     string
	Interest string
}

func (this *Students) GetName() string {
	return this.Name
}

func (this *Students) GetJob() {
	fmt.Println("I'm student, I can swimming and ")
}

func (this *Students) GetSkill() {
	fmt.Println("I'm a student, my skill is play game")
}

func GetHumanInfo(human Humans) {
	fmt.Println(human.GetName())
	human.GetSkill()
	human.GetJob()
}

func main() {

	student := Students{
		Name:     "xiao ming",
		Interest: "swimming",
	}

	teacher := Teacher{
		Name:  "David",
		Skill: "fly",
	}

	GetHumanInfo(&student)

	GetHumanInfo(&teacher)
}
