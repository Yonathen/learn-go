package main

import "fmt"

func forLoop() {
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

func ifCondition(time int) {
	fmt.Println("\nIf Condition")
	if time < 10 {
		fmt.Println("Good morning")
	} else if time > 12 {
		fmt.Println("Good Afternoon")
	} else {
		fmt.Println("Good evening")
	}
}

func switchStatement(day int) {
	fmt.Println("\nSwitch Statement")
	switch day {
	case 1:
		fmt.Println("Sunday")
	case 2:
		fmt.Println("Monday")
	case 3:
		fmt.Println("Tuesday")
	case 4:
		fmt.Println("Wednesday")
	case 5:
		fmt.Println("Thursday")
	case 6:
		fmt.Println("Friday")
	case 7:
		fmt.Println("Saturday")
	default:
		fmt.Println("Invalid day")
	}
}

func main() {
	forLoop()
	ifCondition(3)
	ifCondition(13)
	switchStatement(2)
}
