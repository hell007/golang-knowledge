package stackArray

import (
	stack "Data-Structures/stack/stack-array"
	"errors"
	"fmt"
	"strconv"
	"testing"
)

func TestStack(t *testing.T) {

	s := stack.New()

	// Test IsEmpty
	b := s.IsEmpty()
	fmt.Println("s.IsEmpty==", b)

	// Test Push
	for i := 0; i < 10; i++ {
		err := s.Push("stack-array" + strconv.Itoa(i))
		if err!= nil {
			errors.New("stack push error")
		}
	}

	// Test Len
	len := s.Len()
	fmt.Println("s.Len==", len)

	// Test Top
	m,err := s.Top()
	fmt.Println("s.Top==", m, err)

	// Test Pop
	n,err := s.Pop()
	fmt.Println("s.Pop==", n, err)





}