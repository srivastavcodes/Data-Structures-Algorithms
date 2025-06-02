package main

func twoSum(digits []int, res int) []int {
	record := make(map[int]int)

	for i, dig := range digits {
		sub := res - dig
		if _, ok := record[sub]; ok {
			return []int{record[sub], i}
		}
		record[dig] = i
	}
	return nil
}
