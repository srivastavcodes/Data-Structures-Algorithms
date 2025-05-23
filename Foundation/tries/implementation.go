package main

type TrieNode struct {
	value    byte
	isEnd    bool
	children [26]*TrieNode
}

func NewTrieNode(char byte) TrieNode {
	return TrieNode{value: char}
}

func (root *TrieNode) InsertWord(word string) {
	if len(word) == 0 {
		root.isEnd = true
		return
	}
	idx := word[0] - 'A'

	if root.children[idx] == nil {
		root.children[idx] = &TrieNode{value: word[0]}
	}
	root.children[idx].InsertWord(word[1:])
}

func (root *TrieNode) SearchWord(word string) bool {
	if len(word) == 0 {
		return root.isEnd
	}
	valid := 'A' <= word[0] && word[0] <= 'Z'

	if !valid {
		return false
	}
	idx := word[0] - 'A'

	if root.children[idx] == nil {
		return false
	}
	return root.children[idx].SearchWord(word[1:])
}

func (root *TrieNode) SearchPrefix(word string) bool {
	if len(word) == 0 {
		return true
	}
	valid := 'A' <= word[0] && word[0] <= 'Z'

	if !valid {
		return false
	}
	idx := word[0] - 'A'

	if root.children[idx] == nil {
		return false
	}
	return root.children[idx].SearchPrefix(word[1:])
}
