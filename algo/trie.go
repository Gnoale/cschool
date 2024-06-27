package main

import (
	"fmt"
)

type Trie struct {
	children []*Trie
	end      bool
}

func Constructor() *Trie {
	return &Trie{
		children: make([]*Trie, 26),
		end:      false,
	}
}

/*
	current = [nil, *Trie, nil, Trie, ...]
					|
					v
					[nil, nil, nil, nil, ...]


*/

func (this *Trie) Insert(word string) {
	current := this
	for _, letter := range word {
		idx := alphaToIdx(letter)
		if idx < 0 {
			panic("not alpha")
		}
		//runtime.Breakpoint()
		fmt.Println(idx, string(letter))
		if current.children[idx] == nil {
			current.children[idx] = Constructor()
		}
		current = current.children[idx]
	}
	current.end = true
}

func (this *Trie) startsWith(word string) *Trie {
	current := this
	for _, letter := range word {
		idx := alphaToIdx(letter)
		if idx < 0 {
			panic("not alpha")
		}
		if current.children[idx] == nil {
			return nil
		}
		current = current.children[idx]
	}
	return current
}

func (this *Trie) Search(word string) bool {
	node := this.startsWith(word)
	if node == nil {
		return false
	}
	return node.end
}

func (this *Trie) StartsWith(word string) bool {
	node := this.startsWith(word)
	return node != nil
}

func alphaToIdx(letter rune) int {
	if letter >= 'a' && letter <= 'z' {
		return int(letter - 'a')
	}
	if letter >= 'A' && letter <= 'Z' {
		return int(letter - 'A')
	}
	return -1
}

func main() {
	words := []string{"trie", "insert", "search", "search", "startSwith", "insert", "search"}
	trie := Constructor()
	for _, word := range words {
		trie.Insert(word)
	}

	fmt.Println(trie.StartsWith("Start"))
	fmt.Println(trie.Search("Start"))
	fmt.Println(trie.StartsWith("boat"))

}
