package main

import "fmt"

//入参只能是 [4]int类型，而且还是值拷贝
func printArray(array [4]int) {
	for index, val := range array {
		fmt.Println("index = ", index, "value = ", val)
	}
}

//引用传递,可以改变入参的值，传递的指针
func printSlice(array []int) {
	for index, val := range array {
		fmt.Println("index = ", index, "val = ", val)
	}

	array[0] = 10
}

//几种定义slice方式
func defineSlice() {
	//1. 声明slice是一个切片，并且初始化为1,2,3
	slice1 := []int{1, 2, 3}
	fmt.Println("slice1 length: ", len(slice1), slice1)

	//2.声明slice是一个切片，但是没有给slice分配内存空间
	var slice2 []int
	//length 是0,没有值
	fmt.Println("slice2 length: ", len(slice2), slice2)
	//给slice2开辟两个内存空间, length是2，值都是0
	slice2 = make([]int, 2)
	fmt.Println("after make slice2 length: ", len(slice2), slice2)

	//判断slice是否为空切片
	var emptySlice []int
	if emptySlice == nil {
		fmt.Println("empty slice")
	} else {
		fmt.Println("is not empty slice")
	}
}

//slice的容量(到容量 增加1倍)、长度、截取、拷贝

func sliceDemo() {

	//长度3，容量5
	slice1 := make([]int, 3, 5)
	slice1 = append(slice1, 10)
	fmt.Println("********", "slice1 length = ", len(slice1), ", slice capability = ", cap(slice1), "slice1 = ", slice1)

	slice1[0] = 1
	slice1[1] = 2
	slice1[2] = 3

	// slice1[3] = 4  capability 是5，虽然底层内存空间分配好了，但是长度是3，(3 , 4是不可访问的)

	fmt.Println("slice1: ", slice1) // [1 2 3]

	//可以使用append进行元素追加

	slice1 = append(slice1, 4)
	slice1 = append(slice1, 5)

	fmt.Println("slice1: ", slice1) // [1 2 3 4 5 ]

	fmt.Println("slice1 length = ", len(slice1), ", slice capability = ", cap(slice1)) // 5 , 5

	//###  如果此时cap被站满了，再尝试增加元素，则会动态增加空间，再次开辟cap个空间
	slice1 = append(slice1, 6)

	fmt.Println("slice1: ", slice1) // [1 2 3 4 5 6]

	//#### 此时cap变成10
	fmt.Println("slice1 length = ", len(slice1), ", slice capability = ", cap(slice1)) // 6 , 10

	//####  创建一个slice不指定cap, 此时cap与length相等
	slice2 := make([]int, 3)
	fmt.Println("slice2 length = ", len(slice2), "slice2 cap ", cap(slice2)) //3 , 3

	//### 此时追加一个元素，cap 变成6

	slice2 = append(slice2, 2)
	fmt.Println("slice2 length = ", len(slice2), "slice2 cap ", cap(slice2)) //4 , 6

	//##### 切片截取

	s := []int{1, 2, 3}

	//[0: 2) 左边右开[1, 2]
	s1 := s[0:2]

	fmt.Println(s1)
}

func sliceCopyDemo() {

	slice1 := []int{1, 2, 3}

	fmt.Println("slice1: ", slice1)

	slice2 := slice1[0:2] //[1 2 3]

	//因为slice2和slice1都指向同一块地址，slice2修改，也会修改slice1中的值
	slice2[0] = 99

	fmt.Println("slice2: ", slice2)

	fmt.Println("slice1: ", slice1) //[99 2 3]

	//可以使用copy函数将底层数组做一份拷贝

	slice3 := make([]int, 3)
	copy(slice3, slice1)
	fmt.Println("slice3: ", slice3) //[99 2 3]

	slice3[0] = 9999
	fmt.Println(slice3) // [9999 2 3]
	fmt.Println(slice1) // [99, 2 3]

}
func main() {
	array1 := [4]int{1, 2, 3, 4}

	printArray(array1)

	myArray := []int{1, 2, 3, 4, 5} //切片，动态数组

	fmt.Printf("myArray type is %T\n", myArray)

	printSlice(myArray)

	fmt.Println(myArray)

	fmt.Println("############")
	defineSlice()

	fmt.Println("=============")
	sliceDemo()

	fmt.Println("=========slice copy demo ========")
	sliceCopyDemo()
}
