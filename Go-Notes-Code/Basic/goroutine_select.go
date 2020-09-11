package main

import "fmt"

func fibonacci2(c, quit chan int) {
	fmt.Println("fibonacci2")
	x, y := 1, 1
	for {
		select {
			case c <- x:
				x, y = y, x + y
			case <- quit:
				fmt.Println("quit接收到 0 命令")
				return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		fmt.Println("go func")
		for i := 0; i < 10; i++ {
			m := <- c
			fmt.Println("c接收到的信息=", m)
		}
		quit <- 0
	}()

	fibonacci2(c, quit)
}