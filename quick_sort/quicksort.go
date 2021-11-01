package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	array := []int{16, 4, 10, 14, 7, 9, 3, 2, 8, 1}
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
	pivot := pivot(array, from, to)
	array[to], array[pivot] = array[pivot], array[to]
	return partition(array, from, to)
}

//Изменяет массив так, что слева находятся
//самые маленькие значения, а справа - самые большие
func partition(array []int, from, to int) ([]int, int) {
	//место i-х элементов постепенно занимают элементы,
	//меньшие элемента с индексом to.
	x := array[to]
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

//Выбирает три рандомных элемента из
//массива и вычисляет медиану
func pivot(array []int, from int, to int) int {
	rand.Seed(time.Now().UnixMicro())
	a := rand.Intn(to-from) + from
	b := rand.Intn(to-from) + from
	c := rand.Intn(to-from) + from
	pivot := elementsMedian(array, a, b, c)
	return pivot
}

//Находим медиану трёх элементов и возвращаем индекс
//медианного элемента. Отнюдь не элегантное
//и быстрое решение, однако нам этого хватит сполна т.к. элементов всегда три
func elementsMedian(array []int, a, b, c int) int {
	if array[a] < array[b] {
		switch {
		case array[b] < array[c]:
			return b
		case array[a] < array[c]:
			return c
		default:
			return a
		}
	}
	switch {
	case array[a] < array[c]:
		return a
	case array[b] < array[c]:
		return c
	default:
		return b
	}
}
