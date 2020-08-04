package linkedList

import (
	"fmt"
	"testing"
)

func slice(args ...interface{}) []interface{} {
	return args
}

//go test -v list.go list_test.go
func TestLinkedList(t *testing.T) {

	// test Prepend
	list := NewList()
	fmt.Println("==test Prepend==")
	list.Prepend(NewNode(1))
	list.Prepend(NewNode(2))
	list.Prepend(NewNode(3))
	list.String()

	// Test Append
	fmt.Println("==test Append==")
	k := NewList()
	k.Append(NewNode(4))
	k.Append(NewNode(5))
	k.Append(NewNode(6))
	k.String()

	//zero := *slice(k.Get(0))[0].(*Node).Value.(*Node)
	//fmt.Println(zero.Value)

	fmt.Println("==test Insert==")
	k.Insert(NewNode(7), 1)
	k.String()

	fmt.Println("==test Remove==")
	k.Remove(*NewNode(7))
	k.String()

	fmt.Println("==test Contac==")
	list.Concat(k)
	list.String()

	fmt.Println("==test Find==")
	index, _ := list.Find(NewNode(1))
	fmt.Printf("NewNode(1) 的 index = %d", index)
	list.Get(3)

	fmt.Println()
	fmt.Println("==test Clear==")
	list.Clear()
	list.String()

}