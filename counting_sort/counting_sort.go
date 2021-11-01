package main

import "fmt"

func main() {
	a := []int{16, 4, 10, 100, 14, 7, 9, 3, 2, 8, 1}
	a = countingSort(a)
	fmt.Println(a)
}

func countingSort(array []int) []int {
	k := maxElement(array)

	c := make([]int, k+1)
	for j := 0; j < len(array); j++ {
		c[array[j]]++
	}
	for i := 1; i < k+1; i++ {
		c[i] += c[i-1]
	}
	result := make([]int, len(array)+1)
	for j := len(array) - 1; j > 0; j-- {
		result[c[array[j]]] = array[j]
		c[array[j]]--
	}
	result = result[1:]
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
