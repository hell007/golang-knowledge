package main

import "fmt"

type TrieNode struct {
	Date string
	Children map[string]*TrieNode
	IsEndingChar bool
}

type GoTrie struct {
	Root *TrieNode
}

func (this *GoTrie) Insert(text string) {

	p := this.Root

	for i:=0; i < len(text); i++ {
		index := string(text[i])
		data := index
		if p.Children == nil {
			p.Children = make(map[string]*TrieNode)
		}
		if _, ok := p.Children[index]; !ok {
			newNode := &TrieNode{
				Date:data,
			}
			p.Children[index] = newNode
		}

		p = p.Children[index]
	}

	p.IsEndingChar = true
}

func (this *GoTrie) Find(pattern string) bool {
	p := this.Root

	for i:=0; i<len(pattern); i++ {
		index := string(pattern[i])

		if _, ok := p.Children[index]; !ok {
			return false
		}

		p = p.Children[index]
	}

	if p.IsEndingChar == false {
		return false
	}
	return true
}

func main(){
	strs := [...]string{"Laravel", "Framework", "PHP"}
	trie := &GoTrie{}
	trie.Root = &TrieNode{}

	for _, str := range strs {
		trie.Insert(str)
	}

	fmt.Println(trie.Root.Children["F"].Children)

	fmt.Println(trie.Find("PHP"))
	fmt.Println(trie.Find("PHPA"))
}
