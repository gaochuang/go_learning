package main

import (
	"fmt"
	"reflect"
)

/*
type Kind uint

const (
    Invalid Kind = iota  // 非法类型
    Bool                 // 布尔型
    Int                  // 有符号整型
    Int8                 // 有符号8位整型
    Int16                // 有符号16位整型
    Int32                // 有符号32位整型
    Int64                // 有符号64位整型
    Uint                 // 无符号整型
    Uint8                // 无符号8位整型
    Uint16               // 无符号16位整型
    Uint32               // 无符号32位整型
    Uint64               // 无符号64位整型
    Uintptr              // 指针
    Float32              // 单精度浮点数
    Float64              // 双精度浮点数
    Complex64            // 64位复数类型
    Complex128           // 128位复数类型
    Array                // 数组
    Chan                 // 通道
    Func                 // 函数
    Interface            // 接口
    Map                  // 映射
    Ptr                  // 指针
    Slice                // 切片
    String               // 字符串
    Struct               // 结构体
    UnsafePointer        // 底层指针
)
*/

/*从类型对象中获取类型名称和种类*/
func reflctIntTest() {
	var a int
	typeOfA := reflect.TypeOf(a)
	fmt.Println(typeOfA.Name())        //typeOfA变量的类型名int
	fmt.Println(typeOfA.Kind())        //typeOfA变量的种类int
	fmt.Printf("%T\n", typeOfA.Name()) //返回的是字符串
	fmt.Printf("%T\n", typeOfA.Kind()) //返回的是reflect.Kind 类型的常量
}

type Emum int

const (
	Zero Emum = 0
)

/*从类型对象中获取类型名称和种类*/
func reflectStructTest() {
	type cat struct {
	}

	// 获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(cat{})
	//显示反射类型的对象的名称和种类
	fmt.Printf("name: %s, kind: %v\n", typeOfCat.Name(), typeOfCat.Kind())

	// 获取Zero常量的反射类型对象
	typeOfZeoo := reflect.TypeOf(Zero)
	//显示反射类型的对象的名称和种类
	fmt.Printf("name: %s, kind: %v\n", typeOfZeoo.Name(), typeOfZeoo.Kind())
}

/*指针和指针指向的元素*/
func reflectPrtTest() {
	type cat struct {
	}

	//创建cat实例
	ins := &cat{}
	//获取结构体反射类型
	typeOfCat := reflect.TypeOf(ins)
	//显示反射对象名称和种类,
	//Go语言的反射中对所有指针变量的种类都是 Ptr，但需要注意的是，指针变量的类型名称是空，不是 *cat
	//name: , kind: ptr
	fmt.Printf("name: %s, kind: %v\n", typeOfCat.Name(), typeOfCat.Kind())
	//取类型的元素
	typeOfCat = typeOfCat.Elem()
	//输出指针变量指向元素的类型名称和种类，得到cat的名称和种类
	//name: cat, kind: struct
	fmt.Printf("name: %s, kind: %v\n", typeOfCat.Name(), typeOfCat.Kind())
}

/*
type StructField struct {
    Name string          // 字段名
    PkgPath string       // 字段路径
    Type      Type       // 字段反射类型对象
    Tag       StructTag  // 字段的结构体标签
    Offset    uintptr    // 字段在结构体中的相对偏移
    Index     []int      // Type.FieldByIndex中的返回的索引值
    Anonymous bool       // 是否为匿名字段
}
*/

func reflectGetStructNum() {
	type cat struct {
		Name string
		//带有tag字段
		Type int `json:"type" id:"100"`
	}

	ins := cat{Name: "mini", Type: 1}
	//获取结构体实例的反射型对象
	typeofCat := reflect.TypeOf(ins)
	fmt.Printf("Name: %s, Kind:%v\n", typeofCat.Name(), typeofCat.Kind())
	//遍历结构体中所有的成员变量
	for i := 0; i < typeofCat.NumField(); i++ {
		//获取结构体中每个成员变量的类型和tag
		filedType := typeofCat.Field(i)
		fmt.Printf("filedType name: %v,Tag:%v\n", filedType.Name, filedType.Tag)
	}

	//通过字段查找字段类型信息
	if catType, ok := typeofCat.FieldByName("Type"); ok {
		//从tag中取需要的tag
		fmt.Println(catType.Tag.Get("json"), catType.Tag.Get("id"))
	}

}

/*使用反射获取结构体成员类型*/
func main() {
	reflctIntTest()
	reflectStructTest()
	reflectPrtTest()
	reflectGetStructNum()
}
