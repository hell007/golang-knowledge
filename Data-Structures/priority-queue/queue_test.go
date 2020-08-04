package priorityQueue

import (
	"fmt"
	"testing"
)

func TestMaxPriorityQueue(t *testing.T) {
	h := NewMax()

	// Test IsEmpty
	fmt.Println("h.IsEmpty==", h.IsEmpty())

	// Test Push
	h.Push(*NewItem(8, 10))
	h.Push(*NewItem(7,11))
	h.Push(*NewItem(6, 12))
	h.Push(*NewItem(3, 13))
	h.Push(*NewItem(1, 14))
	h.Push(*NewItem(0, 15))
	h.Push(*NewItem(2, 16))
	h.Push(*NewItem(4, 17))
	h.Push(*NewItem(9, 18))
	h.Push(*NewItem(5, 19))

	// Test Len
	len := h.Len()
	fmt.Println("h.Len==", len)

	fmt.Println("h==", h)
}

func TestMinPriorityQueue(t *testing.T) {
	h := NewMin()

	h.Push(*NewItem(8, 10))
	h.Push(*NewItem(7, 11))
	h.Push(*NewItem(6, 12))
	h.Push(*NewItem(3, 13))
	h.Push(*NewItem(1, 14))
	h.Push(*NewItem(0, 15))
	h.Push(*NewItem(2, 16))
	h.Push(*NewItem(4, 17))
	h.Push(*NewItem(9, 18))
	h.Push(*NewItem(5, 19))

	fmt.Println("h==", h)
}

func TestChangePriority(t *testing.T) {
	h := NewMax()

	h.Push(*NewItem(8, 10))
	h.Push(*NewItem(7, 11))
	h.Push(*NewItem(6, 12))
	h.Push(*NewItem(3, 13))
	h.Push(*NewItem(1, 14))
	h.Push(*NewItem(0, 15))
	h.Push(*NewItem(2, 16))
	h.Push(*NewItem(4, 17))
	h.Push(*NewItem(9, 18))
	h.Push(*NewItem(5, 19))

	h.ChangePriority(8, 66)
	popped := h.Pop()

	fmt.Println("popped==", popped)

	fmt.Println("h==", h)
}