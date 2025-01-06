package main

import "fmt"

func main() {
	fmt.Println("For loop : with continue and break demo")
	for i := 0; i <= 10; i++ {
		if i == 8 {
			break
		}

		if i == 3 {
			continue
		}

		fmt.Printf("Binary : %b %s", i, "\n")
		fmt.Printf("Hex : %x %s", i, "\n")
		fmt.Printf("Decimal : %d %s", i, "\n")
	}

	fmt.Println("\nNested loop")
	adj := [2]string{"Big", "Tasty"}
	fruits := [3]string{"Apple", "Orange", "Banana"}
	for j := 0; j < len(adj); j++ {
		for k := 0; k < len(fruits); k++ {
			fmt.Printf("Kind : %s, Fruit : %s \n", adj[j], fruits[k])
		}
	}

	fmt.Println("\nIterating with Range")
	for index, value := range fruits {
		fmt.Printf("Value %s at index of %d \n", value, index)
	}

	fmt.Println("\nIterating with Range By omitting index")
	for _, value := range fruits {
		fmt.Printf("Value %s \n", value)
	}
}
