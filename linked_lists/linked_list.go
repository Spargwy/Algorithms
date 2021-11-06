package main

import "fmt"

func main() {
	a := []int{16, 4, 10, 14, 7, 9, 3, 2, 8, 1}
	list := New()

	//add elements in original queue
	for i := 0; i < len(a); i++ {
		list.PushTail(a[i])
	}
	PrintList(list)
	fmt.Println()
	el, err := list.DeleteByIndex(8)
	if err != nil {
		fmt.Println(err, el)
	}
	el, err = list.DeleteByIndex(9)
	if err != nil {
		fmt.Println(err, el)
	}

}

func PrintList(list *List) {
	for e := list.Head(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
