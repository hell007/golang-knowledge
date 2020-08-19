package huffman

import (
	"errors"
	"sort"
	"strings"
)

//结点
type Node struct {
	Value  string
	Weight int
	Left   *Node
	Right  *Node
}

//哈夫曼树
type Tree struct {
	Root *Node
}

//通过字符出现的权重返回节点集合
func buildCharByWeight(text string) []Node {
	weights := make(map[string]int)

	for _, char := range text {
		frequency, exists := weights[string(char)]
		if exists {
			weights[string(char)] = frequency + 1
		} else {
			weights[string(char)] = 1
		}
	}

	nodes := []Node{}
	for data, frequency := range weights {
		nodes = append(nodes, Node{Value: data, Weight: frequency})
	}

	return nodes
}

// 节点通过weight排序
func sortNodes(nodes []Node) []Node {
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Weight < nodes[j].Weight
	})
	return nodes
}

func newNode(value string, weight int, left, right *Node) Node {
	return Node{Value:  value, Weight: weight, Left:   left, Right:  right}
}

func copyNode(n *Node) *Node {
	copy := newNode(n.Value, n.Weight, n.Left, n.Right)
	return &copy
}

func removeNode(slice []Node, s int) []Node {
	if len(slice) == 1 {
		return []Node{}
	}
	return append(slice[:s], slice[s+1:]...)
}

// 遍历整个子树
func (n Node) traverse(code string, visit func(string, string)) {
	if leftNode := n.Left; leftNode != nil {
		leftNode.traverse(code+"0", visit)
	} else {
		visit(n.Value, code)
		return
	}
	n.Right.traverse(code+"1", visit)
}

//创建树
func NewHuffmanTree(text string) *Tree {
	nodes := buildCharByWeight(text)

	if len(nodes) < 1 {
		//errors.New("Must contain 2 or more emlments")
		return &Tree{Root: nil}
	}

	for len(nodes) > 1 {
		nodes = sortNodes(nodes)

		newNode := newNode(
			nodes[0].Value+nodes[1].Value,
			nodes[0].Weight+nodes[1].Weight,
			copyNode(&nodes[0]),
			copyNode(&nodes[1]))

		nodes = removeNode(nodes, 0)
		nodes = removeNode(nodes, 0)
		nodes = append(nodes, newNode)
	}

	return &Tree{Root: &nodes[0]}
}

// 加码
func (tree *Tree) Encode(text string) (string, error) {
	if tree.Root == nil {
		return "", errors.New("tree root cannot be nil")
	}

	coding := ""
	for _, ch := range text {
		encoded := ""
		n := tree.Root
		for n.Value != string(ch) {
			if n.Left == nil && n.Right == nil {
				encoded = encoded + "?"
				break
			}
			if n.Left != nil {
				if strings.Contains(n.Left.Value, string(ch)) {
					n = n.Left
					encoded = encoded + "0"
				}
			}
			if n.Right != nil {
				if strings.Contains(n.Right.Value, string(ch)) {
					n = n.Right
					encoded = encoded + "1"
				}
			}
		}
		coding += encoded
	}

	return coding, nil
}

// 解码
func (tree *Tree) Decode(coding string) (string, error) {
	if tree.Root == nil {
		return "", errors.New("tree root cannot be nil")
	}

	n := tree.Root
	decoded := ""
	for i := 0; i < len(coding); i++ {
		ch := string(coding[i])
		if ch == "0" {
			n = n.Left
		}
		if ch == "1" {
			n = n.Right
		}
		if n.Left == nil && n.Right == nil {
			decoded += n.Value
			n = tree.Root
		}
	}

	return decoded, nil
}