/**
 * name: 并发
 * author: jie
 * note:
 *
 * goroutine说到底其实就是协程，但是它比线程更小，十几个goroutine可能体现在底层就是五六个线程，Go语言内部帮你实现了这些goroutine之间的内存共享。执行goroutine只需极少的栈内存(大概是4~5KB)，当然会根据相应的数据伸缩。
 *
 * goroutine比thread更易用、更高效、更轻便。
 *
 * goroutine是通过Go的runtime管理的一个线程管理器。goroutine通过go关键字实现了，其实就是一个普通的函数。
 *
 *
 */

package main

import (
	"fmt"
	"runtime"
	"time"
)

// 1、goroutine简单实用
func print(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

// 总结：我们可以看到go关键字很方便的就实现了并发编程。 上面的多个goroutine运行在同一个进程里面，共享内存数据，不过设计上我们要遵循：不要通过共享来通信，而要通过通信来共享。

// 2、channels

// 那么goroutine之间如何进行数据的通信呢，Go提供了一个很好的通信机制channel.

// 定义一个channel时，也需要定义发送到channel的值的类型。注意，必须使用make 创建channel.

// ch <- v    // 发送v到channel ch.
// v := <-ch  // 从ch中接收数据，并赋值给v

func sum(a []int, c chan int, s string) {
	fmt.Println(s, a)
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total // 发送total到channel c
}

// 3、Buffered Channels
// ch := make(chan type, value)
// 当 value = 0 时，channel 是无缓冲阻塞读写的，当value > 0 时，channel 有缓冲、是非阻塞的，直到写满 value 个元素才阻塞写入。

// ch:= make(chan bool, 4)，创建了可以存储4个元素的bool 型channel。在这个channel 中，前4个元素可以无阻塞的写入。当写入第5个元素时，代码将会阻塞，直到其他goroutine从channel 中读取一些元素，腾出空间

func test() {
	c := make(chan int, 2) //修改2为1就报错,阻塞了，修改2为3可以正常运行，不阻塞。
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}

// 3、Range和Close
// 上面这个例子中，我们需要读取两次c，这样不是很方便，Go考虑到了这一点，所以也可以通过range，像操作slice或者map一样操作缓存类型的channel

func test2(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// for i := range c能够不断的读取channel里面的数据，直到该channel被显式的关闭。上面代码我们看到可以显式的关闭channel，生产者通过内置函数close关闭channel。关闭channel之后就无法再发送任何数据了，在消费方可以通过语法v, ok := <-ch测试channel是否被关闭。如果ok返回false，那么说明channel已经没有任何数据并且已经被关闭。

// 4、Select

// 那么如果存在多个channel的时候，我们该如何操作呢，Go里面提供了一个关键字select，通过select可以监听channel上的数据流动。
// select默认是阻塞的，只有当监听的channel中有发送或接收可以进行时才会运行，当多个channel都准备好的时候，select是随机的选择一个执行的。

func test3(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
			// default:
			// 	return
		}
	}
}

// 在select里面还有default语法，select其实就是类似switch的功能，default就是当监听的channel都没有准备好的时候，默认执行的（select不再阻塞等待channel）

// 6、超时
// 有时候会出现goroutine阻塞的情况，那么我们如何避免整个程序进入阻塞的情况呢？我们可以利用select来设置超时，通过如下的方式实现：

func test4() {
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c:
				println(v)
			case <-time.After(5 * time.Second):
				println("timeout")
				o <- true
				break
			}
		}
	}()
	<-o
}

// 7、runtime goroutine

/*
runtime包中有几个处理goroutine的函数：

Goexit: 退出当前执行的goroutine，但是defer函数还会继续调用

Gosched: 让出当前goroutine的执行权限，调度器安排其他等待的任务运行，并在下次某个时候从该位置恢复执行。

NumCPU: 返回 CPU 核数量

NumGoroutine: 返回正在执行和排队的任务总数

GOMAXPROCS: 用来设置可以并行计算的CPU核数的最大值，并返回之前的值。
*/

func main() {
	// 1.
	go print("world") //开一个新的Goroutines执行
	print("hello")    //当前Goroutines执行

	fmt.Println("")

	// 2.
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)

	go sum(a[:len(a)/2], c, "1 = ") // a[0:3] 17= 7+2+8

	go sum(a[len(a)/2:], c, "2 = ") // a[3:6] -5= -9+4+0

	// 开始执行 2=> x  1=> y

	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

	// 3.
	fmt.Println("")

	test()

	// 4.
	fmt.Println("")

	c1 := make(chan int, 10)
	go test2(cap(c1), c1)
	// for i := range c1能够不断的读取channel里面的数据，直到该channel被显式的关闭
	for i := range c1 {
		fmt.Println(i)
	}

	// 5.
	fmt.Println("")

	c2 := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c2)
		}
		quit <- 0
	}()
	test3(c2, quit)

	// 6.
	fmt.Println("")

	test4()
}
