package stackArray

import (
	"errors"
	"sync"
)

//顺序栈：数组实现
const ARRAY_SIZE = 10

type Stack struct {
	data [ARRAY_SIZE]interface{}
	top int
	lock sync.Mutex
}

func New() *Stack {
	return  &Stack{}
}

// 长度
func (s *Stack) Len() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.top
}

// 空判读
func (s *Stack) IsEmpty() bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.top == 0
}

// 入栈
func (s *Stack) Push(v interface{}) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.top == ARRAY_SIZE {
		return errors.New("The stack is full")
	}

	s.data[s.top] = v
	s.top++
	return nil
}

// 出栈
func (s *Stack) Pop() (interface{}, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.top==0 {
		return 0, errors.New("The stack is empty")
	}
	s.top--
	return s.data[s.top], nil
}

// 栈顶
func (s *Stack) Top() (interface{}, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.top == 0 {
		return 0, errors.New("The stack is empty")
	}

	return s.data[s.top-1], nil
}