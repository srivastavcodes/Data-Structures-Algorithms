package main

func canCompleteCircuit(gas []int, cost []int) int {
	if sum(gas) < sum(cost) {
		return -1
	}
	result, total := 0, 0
	for i := range gas {
		total += gas[i] - cost[i]
		if total < 0 {
			total = 0
			result = i + 1
		}
	}
	return result
}

func sum(nums []int) int {
	var total int
	for _, num := range nums {
		total += num
	}
	return total
}
