package main

import "fmt"

func LinearSearch(Array []int, number int) {
	arraySize := len(Array) - 1

	for i := 0; i <= arraySize; i++ {
		if Array[i] == number {
			fmt.Printf("Your number has index %d", i)
		}
	}
}
