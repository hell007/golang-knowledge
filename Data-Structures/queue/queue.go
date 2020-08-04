package queue

import "sync"

type Queue struct {
	queue []interface{}
	len   int
	lock  *sync.Mutex
}

func New() *Queue{
	return &Queue{
		queue: make([]interface{}, 0),
		len: 0,
		lock: new(sync.Mutex),
	}
}

// 长度
func (q *Queue) Len() int {
	return q.len
}

// 空判断
func (q *Queue) IsEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.len == 0
}

// 入队
func (q *Queue) Push(el interface{}) {
	q.queue = append(q.queue, el)
	q.len++
	return
}

// 出队
func (q *Queue) Shift() (el interface{}) {
	el, q.queue = q.queue[0], q.queue[1:]
	q.len--
	return
}

// 查看 当前即将出队el
func (q *Queue) Peek() (el interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.queue[0]
}