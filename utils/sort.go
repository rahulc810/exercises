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

func Sort() {
	heapSort([]interface{}{3, 52, 6, 83, 1, 45, 90}, true)
}
