package main

func SelectionSort(array []int) {

	lenght := len(array)
	for i := 1; i < lenght; i++ {
		for j := 0; j < lenght; j++ {
			if array[i] < array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}

	}

}
