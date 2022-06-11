package ch9_test

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employ struct {
	Id   string
	Name string
	Age  int
}

func (e *Employ) SetName(name string) {
	e.Name = name
}

func (e *Employ) SetAge(age int) {
	e.Age = age
}

func (e *Employ) SetId(id string) {
	e.Id = id
}

func (e *Employ) GetEmployInfo() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("Id:%s-Name:%s-Age%d", e.Id, e.Name, e.Age)
}

func TestEmploy(t *testing.T) {
	//em := new(Employ)
	var em Employ
	t.Logf("Address is %x\n", unsafe.Pointer(&em.Name))

	em.SetAge(10)
	em.SetId("101")
	em.SetName("Bob")

	t.Log(em.GetEmployInfo())
}
