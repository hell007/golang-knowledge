package stackList

import "container/list"

type Stack struct {
	list *list.List
}

func New() *Stack {
	list := list.New()
	return &Stack{
		list: list,
	}
}

// 长度
func (s *Stack) Len() int {
	return s.list.Len()
}

// 空判读
func (s *Stack) IsEmpty() bool {
	return s.list.Len() == 0
}

// 入栈
func (s *Stack) Push(v interface{}) {
	s.list.PushBack(v)
}

// 出栈
func (s *Stack) Pop() (interface{}, bool) {
	err := s.list.Back()
	if err != nil {
		s.list.Remove(err)
		return err.Value, true
	}
	return nil, false
}

// 获取栈顶
func (s *Stack) Top() (interface{}, bool) {
	e := s.list.Back()
	if e != nil {
		return e.Value, true
	}
	return nil, false
}

