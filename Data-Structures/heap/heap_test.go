package heap

import (
	"fmt"
	"testing"
)

func TestNewMin(t *testing.T) {
	h := NewMin()

	// Test IsEmpty
	fmt.Println("h.IsEmpty==", h.IsEmpty())

	// Test Push
	h.Push(Int(1))
	h.Push(Int(2))

	// Test Len
	len := h.Len()
	fmt.Println("h.Len==", len)

	// Test Pop
	el := h.Pop()
	fmt.Println("h.Pop==", el)

	fmt.Println("h.Len==", h.Len())
}
