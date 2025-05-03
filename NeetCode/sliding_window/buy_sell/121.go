package main

func main() {}

func maxProfit(prices []int) int {
	l, r := 0, 1
	pft := 0
	for r < len(prices) {
		if prices[l] < prices[r] {
			profit := prices[r] - prices[l]
			pft = max(pft, profit)
		} else {
			l = r
		}
		r++
	}
	return pft
}
