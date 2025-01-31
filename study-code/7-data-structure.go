package main

import "fmt"

func practiceArrays() {
	integers := [5]int{1, 2, 3, 4, 5}
	integers3 := [5]int{0: 1, 3: 2, 4: 3, 2: 4, 1: 5}
	languages := [5]string{"Go", "Java", "C++"}
	var integers1 [3]int
	integers2 := [...]int{1, 3, 7, 9, 12}
	integers2D := [2][2]int{{1, 2}, {2, 3}}

	integers1[0] = 1
	integers1[1] = 2
	integers1[2] = 3

	integers2[2] = 8

	fmt.Println(integers)
	fmt.Println(integers1)
	fmt.Println(integers3)
	fmt.Println(len(integers2), integers2[2])
	fmt.Println(languages)

	for _, value := range integers2 {
		fmt.Print(value)
	}

	fmt.Println("\n")
	for _, value := range integers2D {
		for _, item := range value {
			fmt.Print(item)
		}
	}
}

func creatingSlice() {
	numbers := []int{1, 2, 3}
	numbersArray := [8]int{10, 20, 30, 40, 50, 60, 70, 80}

	sliceNumbersFromArray := numbersArray[4:7]

	fmt.Println("numbers", numbers)
	fmt.Println("numbersArray", numbersArray)
	fmt.Println("sliceNumbersArray", sliceNumbersFromArray)
}

func manipulateSlice() {
	sliceNumbers := []int{10, 20, 30, 40, 50, 60, 70, 80}
	sliceNumbersCopy := []int{1, 2, 3}

	fmt.Println("sliceNumbers before append", sliceNumbers)
	sliceNumbers = append(sliceNumbers, 100)
	fmt.Println("sliceNumbers after append", sliceNumbers)

	fmt.Println("sliceNumbersCopy before copy", sliceNumbersCopy)
}

func main() {
	practiceArrays()
	creatingSlice()
}
