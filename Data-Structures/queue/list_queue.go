package queue

import "container/list"

type ListQueue struct {
	mlist *list.List
}

func NewListQueue() *ListQueue {
	l := list.New()
	return &ListQueue{l}
}

// 长度
func (lg *ListQueue) Len() int {
	return lg.mlist.Len()
}

// 空判断
func (lq *ListQueue) IsEmpty() bool {
	return lq.mlist.Len() == 0
}

// 入队：从队尾添加数据到队列
func (lg *ListQueue) Push(el interface{}) {
	lg.mlist.PushBack(el)
}

// 出队或出列: 从队头取出数据
func (lg *ListQueue) Shift() interface{}{
	el := lg.mlist.Front()
	lg.mlist.Remove(el)
	return el.Value
}

// 查看 当前即将出队el
func (lg *ListQueue) Peek() interface{} {
	return lg.mlist.Front().Value
}