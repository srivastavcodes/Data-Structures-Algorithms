package main

func generateParenthesis(count int) []string {
	result := make([]string, 0)
	getParen("", &result, count, 0, 0)
	return result
}

func getParen(str string, res *[]string, count int, left, right int) {
	if left == count && right == count {
		*res = append(*res, str)
		return
	}
	if left < count {
		getParen(str+"(", res, count, left+1, right)
	}
	if right < left {
		getParen(str+")", res, count, left, right+1)
	}
}
