package main

import "fmt"

func main() {
	var student1 string = "Jhon"
	var student2 = "Jane"
	x := 2

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)

	var a, b, c, d int = 1, 3, 5, 7
	var e, f = 6, "Hello world"
	g, h := 7, "Another World"
	var (
		i int
		j int    = 1
		k string = "hello"
	)

	fmt.Println(a, b, c, d)
	fmt.Println(e, f)
	fmt.Println(g, h)
	fmt.Println(i, j, k)
}
