package main

import (
	"fmt"
	"sort"
)

type kv struct {
	Key   int
	Value int
}

var ss []kv

func topKFrequent(nums []int, k int) []int {
	numsMap := make(map[int]int, 0)
	for _, v := range nums {
		numsMap[v]++
	}

	fmt.Printf("%v\n", numsMap)
	for k, v := range numsMap {
		ss = append(ss, kv{k, v})
	}

	fmt.Printf("ss: %v\n", ss)
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	fmt.Printf("ss: %v\n", ss)
	ret := make([]int, 0)

	for _, v := range ss {
		ret = append(ret, v.Key)
	}

	fmt.Printf("ret: %v\n", ret)
	return ret[:k]
}

func main() {
	testNus := []int{-1, -1}

	fmt.Println(topKFrequent(testNus, 1))
}
