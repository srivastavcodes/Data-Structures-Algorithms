package main

import (
	"fmt"
	"sort"
)

func main() {
	position := []int{10, 8, 0, 5, 3}
	speed := []int{2, 4, 1, 1, 3}
	target := 12
	output := carFleet(target, position, speed)
	fmt.Println(output)
}

func carFleet(tgt int, position, speed []int) int {
	pairs := make([][2]int, len(position))

	for i := 0; i < len(position); i++ {
		pairs[i] = [2]int{position[i], speed[i]}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] > pairs[j][0]
	})
	var stack []float64

	for _, pair := range pairs {
		time := float64(tgt-pair[0]) / float64(pair[1])
		stack = append(stack, time)
		if len(stack) >= 2 && stack[len(stack)-1] <= stack[len(stack)-2] {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack)
}
