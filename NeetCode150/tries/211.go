package main

type WordDictionary struct {
	Char     byte
	Terminal bool
	Children [26]*WordDictionary
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (root *WordDictionary) AddWord(word string) {
	if len(word) == 0 {
		root.Terminal = true
		return
	}
	idx := word[0] - 'a'

	if root.Children[idx] == nil {
		root.Children[idx] = &WordDictionary{Char: word[0]}
	}
	root.Children[idx].AddWord(word[1:])
}

func (root *WordDictionary) Search(word string) bool {
	if len(word) == 0 {
		return root.Terminal
	}
	if word[0] == '.' {
		for _, child := range root.Children {
			if child != nil && child.Search(word[1:]) {
				return true
			}
		}
		return false
	} else {
		idx := word[0] - 'a'
		if root.Children[idx] == nil {
			return false
		}
		return root.Children[idx].Search(word[1:])
	}
}
