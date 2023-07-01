package main

import (
	"fmt"
	"sort"
)

type StudentInfo struct {
	name  string
	score int
}

func main() {
	stu := map[string]int{"xiaoMing": 21, "xiaoHong": 20, "xiaoGou": 12, "xiaodong": 99}

	var sortStu []StudentInfo

	for k, v := range stu {
		sortStu = append(sortStu, StudentInfo{k, v})
	}

	fmt.Printf("%v \n", sortStu)

	sort.SliceStable(sortStu, func(i, j int) bool {
		return sortStu[i].score > sortStu[j].score
	})

	fmt.Printf("%v \n", sortStu)

}
