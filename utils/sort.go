package utils

import (
	"fmt"
)

func heapSort(arr []interface{}, inc bool) {
	heap := &HeapArr{arr, 0, len(arr)}
	heap.Heapify()
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i], _ = heap.pop()
	}

	fmt.Println(arr)
}

func merge(arrA, arrB []interface{}) []interface{} {
	//printArr(arr, s, mid)
	//printArr(arr, mid+1, e)

	ret := make([]interface{}, len(arrA)+len(arrB))

	var i, j, c int
	for i < len(arrA) && j < len(arrB) {
		com, _ := Compare(arrA[i], arrB[j])
		if com <= 0 {
			ret[c] = arrA[i]
			i++
		} else {
			ret[c] = arrB[j]
			j++
		}
		c++
	}
	for ; i < len(arrA); i++ {
		ret[c] = arrA[i]
		c++
	}
	for ; j < len(arrB); j++ {
		ret[c] = arrB[j]
		c++
	}

	return ret
}

func mergeSort(arr []interface{}) []interface{} {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	return merge(mergeSort(arr[:mid]), mergeSort(arr[mid:]))

}

func printArr(arr []interface{}, s, e int) {
	for i := s; i <= e; i++ {
		fmt.Printf("%v ", arr[i])
	}
	fmt.Println()
}

func Sort() {

	arr := []interface{}{3, 52, 6, 83, 1, 45, 90}
	//arr = []interface{}{2, 1, 3}
	//heapSort(arr, true)
	arr = mergeSort(arr)
	fmt.Println(arr)
}
