/*
 * @Descripttion:
 * @Author: zenghua.wang
 * @Date: 2021-02-21 21:29:33
 * @LastEditors: zenghua.wang
 * @LastEditTime: 2022-12-08 14:32:56
 */
package stackList

import (
	"fmt"
	stack "golang-knowledge/Data-Structures/stack/stack-list"
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
