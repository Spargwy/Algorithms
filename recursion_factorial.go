package main

func Factorial(factorialNumber uint) uint {
	if factorialNumber == 0 {
		return 1
	}
	factorialOutput := factorialNumber * Factorial(factorialNumber-1)

	return factorialOutput

}
