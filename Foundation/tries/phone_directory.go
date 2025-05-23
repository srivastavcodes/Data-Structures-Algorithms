package main

func (root *TrieNode) getSuggestions(query string) [][]string {
	output := make([][]string, 0)
	prev, prefix := root, ""

	for i := 0; i < len(query); i++ {
		lastChar := query[i]
		prefix = prefix + string(lastChar)

		curr := prev.children[lastChar-'A']
		if curr == nil {
			break
		}
		record := make([]string, 0)

		retrieveSuggestions(curr, &record, prefix)
		output = append(output, record)

		prev = curr
	}
	return output
}

func retrieveSuggestions(curr *TrieNode, record *[]string, prefix string) {
	if curr.isEnd {
		*record = append(*record, prefix)
	}
	for char := 'A'; char <= 'Z'; char++ {
		next := curr.children[char-'A']
		if next != nil {
			prefix2 := prefix + string(char)
			retrieveSuggestions(next, record, prefix2)
		}
	}
}
