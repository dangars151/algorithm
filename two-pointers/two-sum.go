package main

import "fmt"

func main() {
	numbers := []int{3, 24, 50, 79, 88, 150, 345}
	target := 200
	fmt.Println(twoSumUsingTwoPoints(numbers, target))
}

func twoSumUsingTwoPoints(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		sum := numbers[left] + numbers[right]

		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return nil
}