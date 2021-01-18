package main

import "fmt"

func selectionSort(array []int) {

	lenght := len(array)
	for i := 1; i < lenght; i++ {
		for j := 0; j < lenght; j++ {
			if array[i] < array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}

	}
	for el := range array {
		fmt.Printf("%d: %d\n", el, array[el])
	}

}

func main() {
	array := []int{1, 45, 4, 5, 4, 5, 4, 5, 34, 5, 3, 5, 3, 5, 3, 6, 2535646, 3636, 3, 43636, -9000000000, 36363, 524, 536, 3, 532, 63, 6, 36, 2, 62, 6, 26, 2, 62, 6, 26, 2, 6}
	selectionSort(array)
}
