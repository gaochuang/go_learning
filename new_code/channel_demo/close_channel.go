package main

import (
	"fmt"
)

func rangeChannel() {

	ch := make(chan int, 5)
	for i := 0; i < 5; i++ {
		ch <- i
	}

	count := 0
	//range 遍历channel,如果ch没有数据 range会阻塞
	for v := range ch {
		fmt.Println(v)
		if count > 3 {
			break
		}
		count++

	}
}

/*
1.向已经关闭的channel写数据会产生panic
2.关闭channel，可以继续从channel接收数据
3.对于nil channel,无论收发都会被阻塞
*/
func main() {

	ch := make(chan int, 5)

	go func() {

		defer func() {
			err := recover()
			fmt.Println("err: ", err)
		}()

		for i := 0; i < 5; i++ {
			ch <- i
		}
		//关闭channel
		close(ch)

		//产生panic
		ch <- 1
	}()

	for i := 0; i < 6; i++ {

		//ok如果为true表示channel没有关闭，如果为false则表示channel关闭
		if num, ok := <-ch; ok {
			fmt.Println("num = ", num)
		} else {
			break
		}

	}

	rangeChannel()
	fmt.Println("main routine exit")
}
