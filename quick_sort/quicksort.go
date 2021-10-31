package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	array := []int{2, 8, 7, 1, 3, 5, 6, 1}
	array = quickSort(array, 0, len(array)-1)
	fmt.Println(array)
}

func quickSort(array []int, from, to int) []int {
	var q int
	if from < to {
		array, q = randomizedPartition(array, from, to)
		array = quickSort(array, from, q-1)
		array = quickSort(array, q+1, to)
	}
	return array
}

func randomizedPartition(array []int, from, to int) ([]int, int) {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn((to - from))
	array[to], array[i+from] = array[i+from], array[to]
	return partition(array, from, to)
}

//Изменяет массив так, что слева находятся
//самые маленькие значения, а справа - самые большие
func partition(array []int, from, to int) ([]int, int) {
	x := array[to]
	//место i-х элементов постепенно занимают элементы,
	//меньшие элемента с индексом to.
	i := from
	for j := from; j < to; j++ {
		if array[j] < x {
			array[i], array[j] = array[j], array[i]
			i++
		}
	}

	//свап, чтобы якорный элемент изменился
	array[i], array[to] = array[to], array[i]
	return array, i
}
