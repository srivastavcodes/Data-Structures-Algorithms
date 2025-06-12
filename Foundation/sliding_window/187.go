package main

func findRepeatedDnaSequences(str string) []string {
	var result []string
	record := make(map[string]int)

	for i := 0; i <= len(str)-10; i++ {
		current := str[i : i+10]
		if record[current] == 1 {
			result = append(result, current)
		}
		record[current]++
	}
	return result
}
