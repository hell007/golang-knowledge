package main

import (
	"fmt"
	"sync"
)

//WaitGroup 顾名思义，就是用来等待一组操作完成的。WaitGroup 内部实现了一个计数器，用来记录未完成的操作个数，它提供了三个方法：

//Add() 用来添加计数
//Done() 用来在操作结束时调用，使计数减一
//Wait() 用来等待所有的操作结束，即计数变为 0，该函数会在计数不为 0 时等待，在计数为 0 时立即返回

func main(){
	var wg sync.WaitGroup

	wg.Add(2) // 因为有两个动作，所以增加2个计数

	go func() {
		fmt.Println("Goroutine 1")
		wg.Done() // 操作完成，减少一个计数
	}()

	go func() {
		fmt.Println("Goroutine 2")
		wg.Done() // 操作完成，减少一个计数
	}()

	wg.Wait() // 等待，直到计数为0
}