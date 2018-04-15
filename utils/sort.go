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

	ret := make([]interface{}, len(arrA)+len(arrB))

	var i, j, c int
	for i < len(arrA) && j < len(arrB) {
		com := Compare(arrA[i], arrB[j])
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

func QuickSort(arr []interface{}, s, e int) {
	if s >= e {
		return
	}

	p := partition(arr, s, e)
	QuickSort(arr, s, p-1)
	QuickSort(arr, p+1, e)
}

func partition(arr []interface{}, s, e int) int {
	pivot := s
	i, j := s+1, e
	for {
		for ; Compare(arr[i], arr[pivot]) < 0; i++ {
		}
		for ; Compare(arr[j], arr[pivot]) > 0; j-- {
		}
		if i < j {
			swap(arr, i, j)
			i++
			j--
		} else {
			break
		}
	}
	swap(arr, j, pivot)
	return j
}

func printArr(arr []interface{}, s, e int) {
	for i := s; i <= e; i++ {
		fmt.Printf("%v ", arr[i])
	}
	fmt.Println()
}

func Sort() {

	arr := []interface{}{3, 52, 6, 83, 1, 45, 90}
	//arr = []interface{}{54, 26, 93, 17, 77, 31, 44, 55, 20}
	//heapSort(arr, true)
	//arr = mergeSort(arr)
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
