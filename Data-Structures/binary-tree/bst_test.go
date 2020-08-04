package binaryTree

import (
	"fmt"
	"testing"
)

func TestNode_Compare(t *testing.T) {
	n := NewNode(1)
	m := NewNode(2)
	k := NewNode(1)

	fmt.Println("===", n.Compare(m))
	fmt.Println("===", n.Compare(k))
	fmt.Println("===", m.Compare(k))
}

func TestTree(t *testing.T) {
	tree := NewTree(nil)

	//list := [10]int{2, 1, 3, 9, 4, 6, 7, 8, 0, 5}
	list := []int{5}

	for _, v := range list {
		tree.Insert(ElementType(v))
	}

	fmt.Println("Size===", tree.Size)
	fmt.Println("Search===", tree.Search(5).Value)
	fmt.Println("Delete===", tree.Delete(5))
	fmt.Println("Size===", tree.Size)
}
