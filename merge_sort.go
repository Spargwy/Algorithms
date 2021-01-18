package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	slice := []int{1, 2, 5, 0, -3, 231, 8, 16, 4}
	fmt.Println("\nSorted", mergeSort(slice), "\n")
}

func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice

}

func mergeSort(array []int) []int {
	lenght := len(array)
	if lenght == 1 {
		return array
	}

	middle := int(lenght / 2)
	var (
		left  = make([]int, middle)        //создаем два массива каждый из которых равен половине начального массива
		right = make([]int, lenght-middle) //запись именно такая т.к. начальное число может не делиться без остатка
	)

	for i := 0; i < lenght; i++ { //присваиваем каждому элементу этих двух массивов неупорядоченные значения входного массива
		if i < middle { //пока не дошли до середины
			left[i] = array[i]
		} else {
			right[i-middle] = array[i] //после середины присваиваем массиву right
		}
	}
	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))

	i := 0
	/*сравниваем первые элементы массивов. Присваеваем очередному элементу выходного массива значение,
	либо из левой части начального массива, либо из правой.
	Удаляем первый элемент и следовательно дальнейший цикл работает с 2, 3 и так далее элементами*/
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] { // предположим: left = 0-5, 11-15, right = 6-10, 16-20. Цикл присваивает элементы индексам
			result[i] = left[0] //пока левая и правая часть содержат элементы. В данном случае цикл завершается после присваивания 11 элемента левого массива
			left = left[1:]     //result теперь содержит числа от 0 до 15.
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++

	}

	//следующие два цикла присваивают оставшиеся значения, оставшиеся в массивах
	for j := 0; j < len(left); j++ { //так как left у нас пуст, мы пропускаем этот цикл и работаем с right
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ { // если в right массиве остались элементы, значит они были больше или равны
		result[i] = right[j] //по сравнению с последним элементом массива left. Следовательно остальная часть массива result должна быть
		i++                  //заполнена элементами из right.
	}
	fmt.Printf("Hello")
	return
}
