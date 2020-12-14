package main

import "fmt"

// select用于在多个channel上同时进行侦听并收发消息，当任何一个case满足条件时即执行;
// 如果没有可执行的case则会执行default的case;
// 如果没有指定default，则会阻塞程序.

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