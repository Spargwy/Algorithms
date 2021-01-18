package main

import "fmt"

func Factorial(factorialNumber uint) uint {
	if factorialNumber == 0 {
		return 1
	}
	factorialOutput := factorialNumber * Factorial(factorialNumber-1)

	fmt.Printf("number %d factorial equal %d\n", factorialNumber, factorialOutput)
	return factorialOutput

}

func main() {
	var factorialNumber int
	fmt.Scanf("%d\n", &factorialNumber)
	Factorial(uint(factorialNumber))
}
