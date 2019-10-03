package main

import "fmt"

/*
Given an array of integers, return a new array such that each element at index i of the new array is the product of all the numbers in the original array except the one at i.

For example, if our input was [1, 2, 3, 4, 5], the expected output would be [120, 60, 40, 30, 24]. If our input was [3, 2, 1], the expected output would be [2, 3, 6].

Follow-up: what if you can't use division?
*/

/*
Approach:
1.-create structure to keep the product of the left part of the index
2.-create structure to keep the product of the right part of the index
3.- iterate over and multiply index from left and index from right then set on a result structure
*/

func main() {
	numbers := []int{1,2,3,4,5}
	productAllExceptOne(numbers)
}

func productAllExceptOne (numbers []int) []int {
	left := make([]int, len(numbers))
	right:= make([]int, len(numbers))
	results := make([]int, len(numbers))

	//set most lefty to 1
	left[0] =1
	//set most righty to 1
	right[len(numbers)-1] = 1

	for x := 1; x < len(left); x++ {
		left[x] = left[x-1] * numbers[x-1]
	}

	for y := len(right) - 2; y >= 0; y--{
		right[y] = right[y+1] * numbers[y+1]
	}

	for z,_ := range results {
		results[z] = left[z] * right[z]
	}

	/*fmt.Printf("%v\n", numbers)
	fmt.Printf("%v\n", left)
	fmt.Printf("%v\n", right)
	fmt.Printf("%v\n", results)*/

	return results
}
