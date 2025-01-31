package main

import "fmt"

func calculate(num1 int, num2 int) (int, int) {
	addition := num2 + num1
	subtraction := num2 - num1

	return addition, subtraction
}

func factorial(num uint) uint {
	if num == 1 {
		return num
	}

	return num * factorial(num-1)
}

var greet = func() {
	fmt.Println("Hi From Anonymous Func")
}

var multiply = func(num int, num1 int) int {
	return num * num1
}

func main() {
	var num1 int
	var num2 int

	fmt.Println("Enter num1")
	fmt.Scanf("%d", &num1)
	fmt.Println("Enter num2")
	fmt.Scanf("%d", &num2)

	sum, difference := calculate(num1, num2)
	fmt.Printf("\nThe sum and difference of %d and %d are %d and %d respectively\n", num1, num2, sum, difference)

	num2Factorial := factorial(uint(num2))
	fmt.Printf("\nFactorial of %d is %d \n", uint(num2), num2Factorial)
	num1Factorial := factorial(uint(num1))
	fmt.Printf("Factorial of %d is %d \n\n", uint(num1), num1Factorial)

	greet()
	fmt.Printf("Multiply %d and %d using Anonymous Func with Arguments to get %d \n", num2, num1, multiply(num2, num1))
}
