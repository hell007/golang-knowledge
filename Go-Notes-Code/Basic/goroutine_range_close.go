package main

import (
	"fmt"
)

/**
channel的基本操作：

var c chan int   //声明一个int类型的channel，注意，该语句仅声明，不初始化channel
c := make(chan type_name)   //创建一个无缓冲的type_name型的channel，无缓冲的channel当放入1个元素后，后续的输入便会阻塞
c := make(chan type_name, 100)   //创建一个缓冲区大小为100的type_name型的channel

c <- x   //将x发送到channel c中，如果channel缓冲区满，则阻塞当前goroutine
<- c   //从channel c中接收一个值，如果缓冲区为空，则阻塞
x = <- c   //从channel c中接收一个值并存到x中，如果缓冲区为空，则阻塞
x, ok = <- c   //从channel c中接收一个值，如果channel关闭了，那么ok为false（在没有default select语句的前提下），在channel未关闭且为空的情况下，仍然阻塞

close(c)   //关闭channel c
for term := range c {}   //等待并取出channel c中的值，直到channel关闭，会阻塞
*/

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