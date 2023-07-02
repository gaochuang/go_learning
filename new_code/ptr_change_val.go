package main

import "fmt"

func swap(a *int, b *int) {
	var tmp int
	tmp = *a
	*a = *b
	*b = tmp
}

func main() {
	a := 2
	b := 3
	swap(&a, &b)
	fmt.Println(a, b)
}
