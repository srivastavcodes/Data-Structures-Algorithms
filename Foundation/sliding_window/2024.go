package main

func maxConsecutiveAnswers(answerKey string, k int) int {
	result := 0

	// case-1: (F -> T)
	countF, left, right := 0, 0, 0

	for right < len(answerKey) {
		if answerKey[right] == 'F' {
			countF++
		}
		for countF > k {
			if answerKey[left] == 'F' {
				countF--
			}
			left++
		}
		result = max(result, right-left+1)
		right++
	}

	// case-2: (T -> F)
	countT, left, right := 0, 0, 0
	for right < len(answerKey) {
		if answerKey[right] == 'T' {
			countT++
		}
		for countT > k {
			if answerKey[left] == 'T' {
				countT--
			}
			left++
		}
		result = max(result, right-left+1)
		right++
	}
	return result
}
