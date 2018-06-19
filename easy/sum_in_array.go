// Given a list of numbers and a number k, return whether any two numbers from the list add up to k.
// For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.
// Bonus: Can you do this in one pass?

package easy

import (
	"fmt"
)

func SumAvailable(sum int, arr []int) bool {
	lookup := make(map[int]int, len(arr))
	for _, val := range arr {
		other := sum - val
		_, ok := lookup[other]
		if ok {
			return true
		} else {
			lookup[val] = other
		}
	}
	return false
}

func ExecSumInArray() {
	sum := 19
	arr := []int{10, 15, 3, 7}
	fmt.Printf("Does a pair exists in %v which add up to %v ? %v", arr, sum, SumAvailable(sum, arr))
}
