package linkedList

import (
	"errors"
	"fmt"
)

// https://www.jianshu.com/p/21979f91e038 参考地址
type Node struct {
	Value interface{}
	Prev *Node
	Next *Node
}

type List struct {
	Length int //链表的大小
	Head *Node //链表头节点
	Tail *Node //链表尾节点
}

func NewNode(value interface{}) *Node {
	return &Node{Value: value}
}

// 创建链表
func NewList() *List{
	list := new(List)
	list.Length = 0
	return list
}

// 长度
func (list *List) Len() int {
	return list.Length
}

// 空判读
func (list *List) IsEmpty() bool {
	return list.Length == 0
}

// 前置插入
func (list *List) Prepend(value interface{}) {
	node := NewNode(value)

	if list.Len() == 0 {
		list.Head = node
		list.Tail = list.Head
	} else {
		formerHead := list.Head
		formerHead.Prev = node

		node.Next = formerHead
		list.Head = node
	}

	list.Length++
}

// 后添加
func (list *List) Append(value interface{}) {
	node := NewNode(value)

	if list.Len() == 0 {
		list.Head = node
		list.Tail = list.Head
	} else {
		formerTail := list.Tail
		formerTail.Next = node

		node.Prev = formerTail
		list.Tail = node
	}

	list.Length++
}

// 查找
func (list *List) Find(node *Node) (int, error) {
	if list.Len() == 0 {
		return 0, errors.New("Empty list")
	}

	index := 0
	found := -1
	list.Map(func(n *Node){
		index++
		if n.Value == node.Value && found == -1 {
			found = index
		}
	})

	if found == -1 {
		return 0, errors.New("Item not found")
	}

	return found, nil
}

// 获取
func (list *List) Get(index int) (*Node, error) {
	if index > list.Len() {
		return nil, errors.New("Index out of range")
	}

	node := list.Head
	for i:=0; i<index; i++ {
		node = node.Next
	}

	return node, nil
}

// 在索引index处插入一个元素
func (list *List) Insert(value interface{}, index int) error {
	if index > list.Len() {
		return errors.New("index out of range")
	}

	node := NewNode(value)

	// 前置插入
	if list.Length == 0 || index == 0 {
		list.Prepend(value)
		return nil
	}

	// 后添加
	if list.Len()-1 == index {
		list.Append(value)
		return nil
	}

	nextNode, _ := list.Get(index)
	prevNode := nextNode.Prev

	prevNode.Next = node
	node.Prev = prevNode

	nextNode.Prev = node
	node.Next = nextNode

	list.Length++

	return nil
}

// 移除
func (list *List) Remove(value interface{}) error {
	if list.Len() == 0 {
		return errors.New("Empty list")
	}

	if list.Head.Value == value {
		list.Head = list.Head.Next
		list.Length--
		return nil
	}

	found := 0
	for n:=list.Head; n!= nil; n=n.Next {
		if *n.Value.(*Node) == value && found == 0 {
			n.Next.Prev, n.Prev.Next = n.Prev, n.Next
			list.Length--
			found++
		}
	}

	if found == 0 {
		return errors.New("Node not found")
	}

	return nil
}

// 删除
func (list *List) Clear() {
	list.Length = 0
	list.Head = nil
	list.Tail = nil
}

// 合并连接
func (list *List) Concat(k *List) {
	list.Tail.Next, k.Head.Prev = k.Head, list.Tail
	list.Tail = k.Tail
	list.Length += k.Length
}

// 遍历
func (list *List) Each(f func(node Node)) {
	for node := list.Head; node != nil; node = node.Next {
		f(*node)
	}
}

// Map
func (list *List) Map(f func(node *Node)) {
	for node := list.Head; node != nil; node = node.Next {
		n := node.Value.(*Node)
		f(n)
	}
}

//遍历整个链表
func (list *List) String() {
	cur := list.Head
	for cur != nil {
		fmt.Println(cur.Value, "->")
		cur = cur.Next
	}
}