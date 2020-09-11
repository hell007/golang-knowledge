package main

import "fmt"

// goroutine / channel

// 1、多个goroutine之间实现同步
// 这种方式是一种比较完美的解决方案， goroutine / channel 它们也是在 go 里面经常搭配在一起的一对.

// 2、go 里面也提供了更简单的方式 —— 使用 sync.WaitGroup。 参看goroutine_wait_group.go
func main() {
	ch := make(chan struct{})
	count := 2 // count 表示活动的协程个数

	go func() {
		fmt.Println("Goroutine 1")
		ch <- struct{}{} // 协程结束，发出信号
	}()

	go func() {
		fmt.Println("Goroutine 2")
		ch <- struct{}{} // 协程结束，发出信号
	}()

	for range ch {
		// 每次从ch中接收数据，表明一个活动的协程结束
		count--
		fmt.Println("count==", count)
		// 当所有活动的协程都结束时，关闭管道
		if count == 0 {
			close(ch)
		}
	}

}
