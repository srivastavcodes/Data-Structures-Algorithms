package main

func groupAnagrams(strs []string) [][]string {
	res := make(map[[26]int][]string)

	for _, str := range strs {
		var count [26]int
		for _, r := range str {
			count[r-'a']++
		}
		res[count] = append(res[count], str)
	}
	var result [][]string
	for _, grp := range res {
		result = append(result, grp)
	}
	return result
}
