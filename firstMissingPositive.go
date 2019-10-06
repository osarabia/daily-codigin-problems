package main


import "fmt"

/*
Given an array of integers, find the first missing positive integer in linear time and constant space. In other words, find the lowest positive integer that does not exist in the array. The array can contain duplicates and negative numbers as well.

For example, the input [3, 4, -1, 1] should give 2. The input [1, 2, 0] should give 3.

You can modify the input array in-place.
*/

func main() {
	numbers := []int{3, 4, -1, 1}
	missing := firstMissingPositive(numbers)
	fmt.Println(missing)

	numbers = []int{1,2,0}
	missing = firstMissingPositive(numbers)
	fmt.Println(missing)
}

func firstMissingPositive(numbers []int) int {
	var temp int
	for _, v := range numbers {
		if v > 0 && v < len(numbers) {
			temp = numbers[v]
			numbers[v] = v
			if temp > 0 && temp < len(numbers) {
				numbers[temp] = temp
			}
		}
	}


	for x := 1; x < len(numbers);  x++ {
		if x != numbers[x] {
			return x
		}
	}

	return len(numbers)
}
