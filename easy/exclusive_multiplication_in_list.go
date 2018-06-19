//Given an array of integers, return a new array such that each element at index i of the new array is the product of all the numbers in the original array except the one at i.
//For example, if our input was [1, 2, 3, 4, 5], the expected output would be [120, 60, 40, 30, 24]. If our input was [3, 2, 1], the expected output would be [2, 3, 6].
//Follow-up: what if you can't use division?

package easy

import (
	"fmt"
)

func exclusiveProduct(arr []int) []int {
	pre := make([]int, len(arr))
	suf := make([]int, len(arr))

	suffMul := 1
	preMul := 1

	for idx, _ := range arr {
		if idx > 0 {
			pre[idx] = preMul * arr[idx-1]
			preMul = pre[idx]
		}
		if idx < len(arr)-1 {
			suf[len(arr)-2-idx] = suffMul * arr[len(arr)-1-idx]
			suffMul = suf[len(arr)-2-idx]
		}
	}

	pre[0] = 1
	suf[len(arr)-1] = 1

	ret := make([]int, len(arr))
	for idx, _ := range arr {
		ret[idx] = pre[idx] * suf[idx]
	}

	return ret

}

func ExecExclusiveMultiplicationInList() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Printf("%v -> %v", arr, exclusiveProduct(arr))
}
