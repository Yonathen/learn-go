package main

import "fmt"

func scanNameAndAge() {
	fmt.Println("\nScan Name And Age By scan()")
	var name string
	var age int

	fmt.Print("Enter your name and age ")
	fmt.Scan(&name, &age)

	fmt.Printf("\n Hi %s, Your age is %d \n", name, age)

	fmt.Println("\nScan Name And Age By scanln()")
	fmt.Print("Enter your name and age ")
	fmt.Scanln(&name, &age)
	fmt.Printf("\n Hi again %s, this is scanln. Your age is %d \n", name, age)

}
func main() {
	scanNameAndAge()
}
