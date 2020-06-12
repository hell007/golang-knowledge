package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x //发送x到channel c.
		x, y = y, x + y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	//_, ok := <-c
	//fmt.Println(ok)
	for i := range c {
		fmt.Println(i)
	}
}