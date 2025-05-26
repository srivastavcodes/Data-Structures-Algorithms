package main

func minimumDays(days int, consume int, allowed int) int {
	var sundays = days / 7
	buyingDays := days - sundays

	totalFood := days * consume
	var ans int

	if totalFood%allowed == 0 {
		ans = totalFood / allowed
	} else {
		ans = totalFood/allowed + 1
	}
	if ans <= buyingDays {
		return ans
	}
	return -1
}
