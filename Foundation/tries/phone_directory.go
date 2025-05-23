package main

func (root *Trie) getSuggestions(query string) [][]string {
	output := make([][]string, 0)
	prev, prefix := root, ""

	for i := 0; i < len(query); i++ {
		lastChar := query[i]
		prefix = prefix + string(lastChar)

		curr := prev.Children[lastChar-'A']
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

func retrieveSuggestions(curr *Trie, record *[]string, prefix string) {
	if curr.IsEnd {
		*record = append(*record, prefix)
	}
	for char := 'A'; char <= 'Z'; char++ {
		next := curr.Children[char-'A']
		if next != nil {
			prefix2 := prefix + string(char)
			retrieveSuggestions(next, record, prefix2)
		}
	}
}
