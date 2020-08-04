package queue

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	queue := New()

	// Test IsEmpty
	fmt.Println("queue.IsEmpty==", queue.IsEmpty())

	// Test Push
	queue.Push(1)
	queue.Push("peek看见我")
	queue.Push(2)
	queue.Push(3)

	// Test Len
	len := queue.Len()
	fmt.Println("queue.Len==", len)

	// Test Shift
	el := queue.Shift()
	fmt.Println("queue.Shift==", el)

	// Test Peek
	fmt.Println("queue.Peek==", queue.Peek())
}


func TestNewListQueue(t *testing.T) {
	queue := NewListQueue()
	fmt.Println("queue.IsEmpty==", queue.IsEmpty())

	// Test Push
	queue.Push("a")
	queue.Push("peek看见我")
	queue.Push("c")
	queue.Push("d")

	// Test Len
	len := queue.Len()
	fmt.Println("queue.Len==", len)

	// Test Shift
	el := queue.Shift()
	fmt.Println("queue.Shift==", el)

	// Test Peek
	fmt.Println("queue.Peek==", queue.Peek())
}