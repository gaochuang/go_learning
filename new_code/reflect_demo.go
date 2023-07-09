/*
                变量
              /    \
             type   value      pair
            /        \
           /          \
          /            \
    static type ||
    concrete type

static type: 在编码中能看的到的类型(int, string)
concrete type: 是runtime系统看见的类型
*/

package main

import (
	"fmt"
	"reflect"
)

func demo1() {
	var a float64 = 3.1415

	//reflect.TypeOf可以得到想要的type类型，例如float64、int32...
	fmt.Println("type: ", reflect.TypeOf(a))
	//reflect.ValueOf可以得到想要的具体值
	fmt.Println("value: ", reflect.ValueOf(a))
}

type User struct {
	Id   int
	Name string
	Age  int
}

func (this User) GetName() string {
	return this.Name
}

func (this User) GetId() int {
	return this.Id
}

/*

这种形式的通过reflect.Type.NumMethod获取不到
func (this *User) GetAge() int {
    return this.Age
}
*/

/*

这种形式的通过reflect.Type.NumMethod也获取不到
func (this *User) getAge() int {
    return this.Age
}
*/

func (this User) GetAge() int {
	return this.Age
}

func demo2(input interface{}) {
	//获取Input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("input type: ", inputType.Name()) //User

	//获取Input的value
	inPutValue := reflect.ValueOf(input)
	fmt.Println("input value: ", inPutValue) //{1 Bob 12}

	//通过type获取struct里面字段
	//1. 获取interface的reflect.Type,通过Type得到NumField，进行遍历
	//2. 得到field，数据类型
	//3. 通过field有一个Interface()方法得到对应的Value

	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		val := inPutValue.Field(i).Interface()

		fmt.Println("field Name: ", field.Name, ", field type: ", field.Type, ", field val: ", val)
	}

	//通过type获取struct的方法
	fmt.Println("numbers of method: ", inputType.NumMethod())
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Println(m.Name, ": ", m.Type)
	}

}

func main() {
	demo1()

	user := User{
		Id:   1,
		Name: "Bob",
		Age:  12,
	}
	demo2(user)
	fmt.Println(user.GetName())
}
