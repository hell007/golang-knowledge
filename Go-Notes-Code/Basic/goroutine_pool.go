package main

import (
	"fmt"
	"sync"
)

// Pool Goroutine Pool
type Pool struct {
	queue chan int
	wg    *sync.WaitGroup
}

// New 新建一个协程池
func New(size int) *Pool {
	if size <= 0 {
		size = 1
	}
	return &Pool{
		queue: make(chan int, size),
		wg:    &sync.WaitGroup{},
	}
}

// Add 新增一个执行
func (p *Pool) Add(delta int) {
	fmt.Println("delta===", delta)
	// delta为正数就添加
	for i := 0; i < delta; i++ {
		p.queue <- 1
	}
	// delta为负数就减少
	for i := 0; i > delta; i-- {
		<-p.queue
	}
	p.wg.Add(delta)
}

// Done 执行完成减一
func (p *Pool) Done() {
	<-p.queue
	p.wg.Done()
}

// Wait 等待Goroutine执行完毕
func (p *Pool) Wait() {
	p.wg.Wait()
}

func main() {
	// 这里限制100个并发
	pool := New(10)// sync.WaitGroup{}

	//假设需要发送1000万个http请求，然后我并发100个协程取完成这件事
	for i := 0; i < 1000; i++ {
		pool.Add(1)  //发现已存在100个人正在发了，那么就会卡住，直到有人完成了宣布自己退出协程了
		go func(i int) {
			//resp, err := http.Get("http://ip.3322.org")
			//if err != nil {
			//	fmt.Println("===>", i, err)
			//} else {
			//	defer resp.Body.Close()
			//	result, _ := ioutil.ReadAll(resp.Body)
			//	fmt.Println("--->", i, string(result))
			//}
			fmt.Println("--->", i)
			pool.Done()
		}(i)
	}
	pool.Wait()
}