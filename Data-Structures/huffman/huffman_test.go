package huffman

import (
	"fmt"
	"testing"
)

func TestNewHuffmanTree(t *testing.T) {
	text := "AAAAABBBBBBBCCDDDDDDDDDDDDD"//A5B7C2D13
	tree := NewHuffmanTree(text)
	fmt.Println("TestNewHuffmanTree==", tree.Root)
}

func TestEncode(t *testing.T) {
	text := "AAAAABBBBBBBCCDDDDDDDDDDDDD"//A5B7C2D13
	tree := NewHuffmanTree(text)
	coding,err := tree.Encode(text)
	if err == nil {
		fmt.Println("coding==", coding, err) //111111111111111101010101010101101100000000000000
	}
}

func TestDecode(t *testing.T) {
	text := "111223"
	tree := NewHuffmanTree(text)
	coding,err1 := tree.Encode(text)
	fmt.Println("coding==", coding, err1)

	str,err2 := tree.Decode(coding)
	fmt.Println("str==", str, err2)
}
