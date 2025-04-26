package main

import "math"

func asteroidCollision(asteroids []int) []int {
	var stack []int
	for i := 0; i < len(asteroids); i++ {
		if asteroids[i] > 0 {
			stack = append(stack, asteroids[i])
		} else {
			for len(stack) > 0 && stack[len(stack)-1] > 0 &&
				stack[len(stack)-1] < int(math.Abs(float64(asteroids[i]))) {
				stack = stack[:len(stack)-1]
			}
			if len(stack) > 0 &&
				stack[len(stack)-1] == int(math.Abs(float64(asteroids[i]))) {
				stack = stack[:len(stack)-1]
			} else if len(stack) == 0 || stack[len(stack)-1] < 0 {
				stack = append(stack, asteroids[i])
			}
		}
	}
	return stack
}
