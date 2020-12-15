package main

import (
	"fmt"
	"strconv"
	"sync"
)

/**
go实现协程池管理
*/

type Productor struct {
	Jobs chan *Job
}

type Consumer struct {
	PoolSize int
	Handler func(chan *Job)(b bool)
	Wg sync.WaitGroup
}

// 任务对象
type Task struct {
	Productor
	Consumer
}

// 任务数据对象
type Job struct {
	Data string
}

// 创建任务
func NewTask(handler func(jobs chan *Job)(b bool))(t *Task){
	t = &Task{
		Productor:Productor{Jobs: make(chan *Job, 100)},
		Consumer:Consumer{PoolSize:100, Handler:handler},
	}
	return
}

// 设置消费者size
func(t *Task)setConsumerPoolSize(poolSize int){
	t.Productor.Jobs = make(chan *Job, poolSize * 10)
	t.Consumer.PoolSize = poolSize
}

// 新增数据
func (c Productor)AddData(data *Job){
	c.Jobs <- data
}

//异步开启多个work去处理任务，但是所有work执行完毕才会退出程序
func (c Consumer)disposeData(data chan *Job){
	//for i:=0; i<=c.PoolSize; i++{
	c.Wg.Add(c.PoolSize)
	go func() {
		defer func() {
			c.Wg.Done()
		}()
		c.Handler(data)
	}()
	//}
	c.Wg.Wait()
}




func main(){
	//1.先实现一个用于处理数据的闭包，在这里面实现自己业务
	consumerHandler := func(jobs chan *Job)(b bool) {
		for job := range jobs {
			fmt.Println(job)
		}
		return
	}

	//2.new一个任务处理对象出来
	t :=NewTask(consumerHandler)
	t.setConsumerPoolSize(100)//500个协程同时消费

	//3.根据自己的业务去生产数据通过AddData方法去添加数据到生产channel,这里是100万条数据
	go func(){
		for i := 0; i < 1000000; i++ {
			job := new(Job)
			iStr := strconv.Itoa(i)
			job.Data = "这里面去定义你的任务数据格式"+ iStr
			t.AddData(job)
		}
	}()

	//4.消费者消费数据
	t.Consumer.disposeData(t.Productor.Jobs)
}
