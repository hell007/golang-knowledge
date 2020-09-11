package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	fmt.Println("fibonacci")
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x //发送x到channel c.
		x, y = y, x + y
	}
	fmt.Println("close")
	close(c)
}

func main() {
	c := make(chan int, 10)
	fmt.Println("c==", c)
	go fibonacci(cap(c), c)

	//v, ok := <-c
	//fmt.Println(v, ok)

	for i := range c {
		fmt.Println(i)
	}
}