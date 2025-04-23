package main

func permutations(res, original string) []string {
	if len(original) == 0 {
		return []string{res}
	}
	ch := original[0]

	var list []string
	for i := 0; i <= len(res); i++ {
		p := res[0:i]
		s := res[i:]
		list = append(permutations(p+string(ch)+s, original[1:]), list...)
	}
	return list
}
