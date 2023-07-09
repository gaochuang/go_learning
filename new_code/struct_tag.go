package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}

func getTag(input interface{}) {
	t := reflect.TypeOf(input).Elem() //Elem获取这个指针指向的元素类型

	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("json")
		fmt.Println("json: ", tag)
	}
}

func main() {

	stu := student{Name: "Bob", Age: 12, Sex: "male"}
	getTag(&stu)
}
