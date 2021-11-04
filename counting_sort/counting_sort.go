package main

import "fmt"

func main() {
	a := []int{16, 4, 10, 100, 14, 7, 9, 3, 2, 0, 8, 1}
	a = countingSort(a)
	fmt.Println(a)
}

func countingSort(arr []int) []int {

	k := maxElement(arr)
	count := make([]int, k+1)

	for i := 0; i < len(arr); i++ {
		count[arr[i]] = count[arr[i]] + 1
	}

	for i := 1; i < k+1; i++ {
		count[i] = count[i] + count[i-1]
	}

	result := make([]int, len(arr)+1)
	for j := 0; j < len(arr); j++ {
		result[count[arr[j]]] = arr[j]
		count[arr[j]] = count[arr[j]] - 1
	}

	return result
}
func maxElement(array []int) int {
	max := array[0]
	for i := 0; i < len(array); i++ {
		if array[i] > max {
			max = array[i]
		}
	}
	return max
}
