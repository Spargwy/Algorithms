package main

import "fmt"

func insertionSort(array []int) {
	/*Пока i меньше длины массива
	присваиваем j значение i для следующего цикла
	значение изначально равно 1, с каждым проходом цикла, сортируются соседние значения. После первого цикла
	j = 2 и цикл повторяется уже два раза, чтобы удостовериться, что ранее отсортированные элементы
	останутся на нужных местах.	*/
	lenght := len(array) //размер массива

	for i := lenght - 1; i > 0; i-- {
		j := i
		for j < lenght { //
			if array[j-1] < array[j] {
				array[j-1], array[j] = array[j], array[j-1]
			}
			j++
		}
	}
	for el := range array {
		fmt.Println(array[el])
	}
}
