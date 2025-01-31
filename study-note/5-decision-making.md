# Decision-Making

## If Statement
- Go supports the usual comparison operators
  - `<, >, <=, >=, ==, and !=`
- Logic operators
  - `&&, ||, !`
- Syntax
  ```
    if conditon_1 {
        // Code block
    } else if condition_2 {
        // Code block
    } else {
        // Code block
    }
  ```

- Example

  ```go
    package main

    import "fmt"

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

    func main() {
        ifCondition(3)
        ifCondition(13)
    }

  ```

## Switch Statement
- The `switch` statement in go runs only the matched case 
  - Therefor it does not need `break` statement
  - Syntax
    ```
        switch expression {
            case x:
                // Code block
            case y,z:
                // code block
            default:
                // Code block
        }
    ```

## Loop

### While
- Syntax
    ```
       for condition {
            // Code Block
        }  
    ```
- Example
```go
package main

import "fmt"

func main()  {
    i := 0
	for i < 10 {
		fmt.Printf("i = %d", i)
		i++
    }
}
```

### do...while
```go
package main

import "fmt"

func main()  {
    i := 0
	for {
		if i > 10 {
			break
        }
		
		fmt.Printf("i = %d", i)
		i++
    }
}
```

### for loop
- Syntax for `for` loop
    ```
    // Syntax for for loop
    for <statment1>; statment2; statment3 {
        // Code block
    } 
    ```
    - `statement1` : initializes the loop
      - `statement2` : Evaluates each loop
      - `staement3` : increase the loop counter value

- Syntax for `for` loop with Range Keyword
    ```
    // Syntax for for loop with Range Keyword
    for <index>, <value> := range <array|slice|map> {
        // Code block
    }
    ```
    - if you want to skip the index use `_` in place of `<index>`

- Example
```go
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
```
