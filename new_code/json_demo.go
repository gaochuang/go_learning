package main

import (
	"encoding/json"
	"fmt"
)

//成员首字母必须大写，才能被json序列化

type Movie struct {
	Title  string   `json:"title"`
	Price  int      `json:"price"`
	Actors []string `json:"actors"`
}

func main() {

	movie := Movie{Title: "one world, one dream", Price: 10, Actors: []string{"Bob", "David", "Jack"}}

	marshal, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	fmt.Printf("json string: %s\n", marshal)

	getMov := &Movie{}

	if err := json.Unmarshal(marshal, getMov); err != nil {
		fmt.Println("unmarshal failed")
		return
	}

	fmt.Printf("%+v\n", *getMov)

}
