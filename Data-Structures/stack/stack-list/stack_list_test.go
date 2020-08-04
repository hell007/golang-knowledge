package stackList

import (
	stack "Data-Structures/stack/stack-list"
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {

	s := stack.New()

	// Test IsEmpty
	b := s.IsEmpty()
	fmt.Println("s.IsEmpty==", b)

	// Test push
	s.Push("111")
	s.Push("222")

	// Test Len
	len := s.Len()
	fmt.Println("s.Len==", len)

	// Test Top
	v1, ok1 := s.Top()
	fmt.Println("s.Top==", v1, ok1)

	// Test pop
	v, ok := s.Pop()
	fmt.Println("s.Pop==", v, ok)

	v, ok = s.Pop()
	fmt.Println("s.Pop==", v, ok)

	v, ok = s.Pop()
	fmt.Println("s.Pop==", v, ok)
}