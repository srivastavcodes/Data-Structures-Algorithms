package main

type Trie struct {
	Value    byte
	IsEnd    bool
	Children [26]*Trie
}

func NewTrieNode(char byte) Trie {
	return Trie{Value: char}
}

func (root *Trie) InsertWord(word string) {
	if len(word) == 0 {
		root.IsEnd = true
		return
	}
	idx := word[0] - 'A'

	if root.Children[idx] == nil {
		root.Children[idx] = &Trie{Value: word[0]}
	}
	root.Children[idx].InsertWord(word[1:])
}

func (root *Trie) SearchWord(word string) bool {
	if len(word) == 0 {
		return root.IsEnd
	}
	valid := 'A' <= word[0] && word[0] <= 'Z'

	if !valid {
		return false
	}
	idx := word[0] - 'A'

	if root.Children[idx] == nil {
		return false
	}
	return root.Children[idx].SearchWord(word[1:])
}

func (root *Trie) SearchPrefix(word string) bool {
	if len(word) == 0 {
		return true
	}
	valid := 'A' <= word[0] && word[0] <= 'Z'

	if !valid {
		return false
	}
	idx := word[0] - 'A'

	if root.Children[idx] == nil {
		return false
	}
	return root.Children[idx].SearchPrefix(word[1:])
}
