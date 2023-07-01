package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func main() {
	f, err := os.Create("./trace.out")
	if err != nil {
		panic(err.Error())
	}

	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("hell trace")

	trace.Stop()

}
