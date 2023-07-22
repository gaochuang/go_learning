package main

import "fmt"

//有缓存的channel

/*
1.当channel已经满了，再向channel写数据会阻塞
2.当channel为空，取数据时，会发生阻塞
*/
func main() {
	ch := make(chan int, 3)

	go func() {
		defer fmt.Println("go routine end")
		for i := 0; i < 4; i++ {
			ch <- i
			fmt.Println("go routine running, send = ", i, "len(ch) = ", len(ch), "cap(ch) = ", cap(ch))
		}
	}()

	for i := 0; i < 4; i++ {
		num := <-ch
		fmt.Println("num = ", num)
	}

	fmt.Println("main routine end")
}
