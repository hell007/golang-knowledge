package priorityQueue

import (
	"Data-Structures/heap"
	"Data-Structures/queue"
	"errors"
)

type Item struct {
	Value    interface{}
	Priority int
}

func NewItem(value interface{}, priority int) (el *Item) {
	return &Item{
		Value:    value,
		Priority: priority,
	}
}

func (x Item) Less(than heap.Item) bool {
	return x.Priority < than.(Item).Priority
}

type PriorityQueue struct {
	data heap.Heap
}

// 最大堆
func NewMax() (q *PriorityQueue) {
	return &PriorityQueue{
		data: *heap.NewMax(),
	}
}

// 最小堆
func NewMin() (q *PriorityQueue) {
	return &PriorityQueue{
		data: *heap.NewMin(),
	}
}

// 长度
func (q *PriorityQueue) Len() int {
	return q.data.Len()
}

// 空判断
func (q *PriorityQueue) IsEmpty() bool {
	return q.data.IsEmpty()
}

// 入队
func (q *PriorityQueue) Push(el Item) {
	q.data.Push(heap.Item(el))
}

// 出队
func (q *PriorityQueue) Pop() Item {
	return q.data.Pop().(Item)
}

// 更改优先级
func (q *PriorityQueue) ChangePriority(v interface{}, priority int) {
	var storage = queue.New()
	//var storage = queue.NewListQueue()

	peek := q.Pop()

	for v != peek.Value {
		if q.Len() == 0 {
			errors.New("Item not found")
		}
		storage.Push(peek)
		peek = q.Pop()
	}

	peek.Priority = priority
	q.data.Push(peek)

	for storage.Len() > 0 {
		q.data.Push(storage.Shift().(heap.Item))
	}
}



















