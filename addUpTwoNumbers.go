package main

import "fmt"

/*
Given a list of numbers and a number k, return whether any two numbers from the list add up to k.

For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.
Bonus: Can you do this in one pass?
*/

func main() {
	numbers1, k1 := []int{10, 15, 3, 7}, 17
	result := addUpNumber(numbers1, k1)
	fmt.Printf("numbers:%v,k:%v,%v\n", numbers1, k1, result)

        numbers2, k2 := []int{10, 15, 3}, 17
	result = addUpNumber(numbers2, k2)
	fmt.Printf("numbers:%v,k:%v,%v\n", numbers2, k2, result)

}

func addUpNumber(numbers []int, k int) bool {
	//create structure to keep numbers already seen
	var remaining int
	seen := make(map[int]bool)

	for _,n := range numbers {
		remaining =  k - n
		// check if we already saw the remaining on our structure
		_, v := seen[remaining]
		if !v {
			seen[n] = true
		} else {
			return true
		}
	}

	return false
}
