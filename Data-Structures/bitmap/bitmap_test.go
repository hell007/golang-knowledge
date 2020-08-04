package bitmap

import (
	"fmt"
	"testing"
)

func TestNewBitMap(t *testing.T) {
	max := 100
	b := NewBitMap(max)
	fmt.Println("bitmap==", b)
	fmt.Println("bitmap String==", b.String())

	// Test Add
	b.Add(13)
	b.Add(100)
	b.Add(0)

	// Test max
	m := b.Max()
	fmt.Println("b.Max==", m)

	// Test IsExist
	bool := b.IsExist(13)
	fmt.Println("b.IsExist==", bool)

	// Test Remove
	b.Remove(13)
	fmt.Println("b.IsExist==", b.IsExist(13))
}