package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine will end")
		fmt.Println("goroutine is running")

		c <- 666
	}()

	num := <-c //从c channel读取数据

	fmt.Println("num = ", num)

	fmt.Println("main goroutine end")
}
