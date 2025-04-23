package main

func subseq(res, original string, list *[]string) {
	if len(original) == 0 {
		*list = append(*list, res)
		return
	}
	ch := original[0]
	subseq(res+string(ch), original[1:], list)
	subseq(res, original[1:], list)
}

func subseq2(res, original string) []string {
	if len(original) == 0 {
		return []string{res}
	}
	ch := original[0]

	left := subseq2(res+string(ch), original[1:])
	right := subseq2(res, original[1:])

	return append(left, right...)
}
