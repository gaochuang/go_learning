package ch27_test

import (
	"reflect"
	"testing"
)

type Employee struct {
	EmployeeId string
	Name       string `format:"normal"`
	Age        int
}

func (e *Employee) UpdateAge(newVal int) {
	e.Age = newVal
}

/*
reflect.Typeof和reflect.Valyeof都有FieldByName方法，但是需要注意他们的区别
*/

func TestInvokeByName(t *testing.T) {
	e := &Employee{"1", "Mike", 30}
	t.Logf("Name: value(%[1]v),Type(%[1]T)", reflect.ValueOf(*e).FieldByName("Name"))

	if NameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		t.Error("get name field failed")
	} else {
		t.Log("Tag:format ", NameField.Tag.Get("format"))
	}

	reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(99)})

	t.Log("Update Age: ", e)

}
