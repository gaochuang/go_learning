package main

import "fmt"

//入参是一个引用传递，值可以修改
func changeVal(input map[int]string) {
	input[1] = "ont + one"
}

//入参是一个引用传递
func printMap(input map[int]string) {

	fmt.Println("print map func")
	fmt.Println(input)
}

func mapDemo() {
	myMap := make(map[int]string)

	myMap[1] = "one"
	myMap[2] = "two"
	myMap[3] = "three"
	myMap[4] = "four"

	fmt.Println(myMap)

	printMap(myMap)

	changeVal(myMap)

	fmt.Println("after change")
	printMap(myMap)

	delete(myMap, 3)
	fmt.Println("after delete")
	printMap(myMap)

}

func main() {
	//声明myMap1是一种map类型，key是string, value是string

	var myMap map[string]string

	if myMap == nil {
		fmt.Println("myMap is null")
	}
	//在使用map前，需要先用make给map分配空间
	myMap = make(map[string]string, 10)

	myMap["one"] = "java"
	myMap["two"] = "c++"
	myMap["three"] = "python"

	fmt.Println(myMap)

	myMap2 := make(map[string]string)
	myMap2["1"] = "one"
	myMap2["2"] = "two"
	myMap2["3"] = "three"

	fmt.Println(myMap2)

	myMap3 := map[string]string{
		"1": "one",
		"2": "two",
		"3": "three",
	}

	fmt.Println(myMap3)

	fmt.Println("================")
	mapDemo()
}
