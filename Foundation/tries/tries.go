package main

import "fmt"

func main() {
	trie := NewTrieNode(0)

	trie.InsertWord("APPLE")
	trie.InsertWord("AMAZON")
	trie.InsertWord("AMERICA")
	trie.InsertWord("BALL")
	trie.InsertWord("BAT")
	trie.InsertWord("BATMAN")

	// Demo the phone directory functionality
	fmt.Println("Phone Directory Suggestions:")

	// Get suggestions for "A"
	fmt.Println("\nSuggestions for 'A':")
	suggestions := trie.getSuggestions("A")
	for i, words := range suggestions {
		fmt.Printf("After typing %d character(s): %v\n", i+1, words)
	}

	// Get suggestions for "AM"
	fmt.Println("\nSuggestions for 'AM':")
	suggestions = trie.getSuggestions("AM")
	for i, words := range suggestions {
		fmt.Printf("After typing %d character(s): %v\n", i+1, words)
	}

	// Get suggestions for "BA"
	fmt.Println("\nSuggestions for 'BA':")
	suggestions = trie.getSuggestions("BA")
	for i, words := range suggestions {
		fmt.Printf("After typing %d character(s): %v\n", i+1, words)
	}

	// Get suggestions for a non-existent prefix
	fmt.Println("\nSuggestions for 'CAR':")
	suggestions = trie.getSuggestions("CAR")
	if len(suggestions) == 0 {
		fmt.Println("No suggestions found")
	} else {
		for i, words := range suggestions {
			fmt.Printf("After typing %d character(s): %v\n", i+1, words)
		}
	}
}
