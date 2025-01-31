# Go Functions

- is a block of statements that can be used repeatedly
- will not execute automatically
- executed by a call to the function

## Creating a function
- Use a `func` keyword
- Syntax
  ```
  func <function_name>(param1 <data_type1>, param2 <data_type2)(<return_type1, return_type2>) {
    // Code block
    return value1, value2
  }
  ```
  
- Example
```go
package main

import "fmt"

func calculate(num1 int, num2 int) (int, int) {
  addition := num2 + num1
  subtraction := num2 - num1

  return addition, subtraction
}

func main() {
  var num1 int
  var num2 int

  fmt.Println("Enter num1")
  fmt.Scanf("%d", &num1)
  fmt.Println("Enter num2")
  fmt.Scanf("%d", &num2)

  sum, difference := calculate(num1, num2)
  fmt.Printf("The sum and difference of %d and %d are %d and %d respectively\n", num1, num2, sum, difference)
}
```

## Variable scope
- Local Variables : Declared within a function
- Global Variables : Declared outside a function

## Anonymous Functions
- These are functions without a function name
```go
package main

import "fmt"

var greet = func () {
	fmt.Println("Hi From Anonymous Func")
}

func main()  {
  greet()
}
```

### Anonymous Functions With Parameters
```go
package main

import "fmt"

var multiply = func(num, num1 int) int {
	return num * num1
}

func main() {
	var num1 int = 10
	var num2 int = 45
	
	fmt.Println(num2, num1, multiply(num2, num1))
}
```

### Return Value from Anonymous Function
```go
package main

import "fmt"

var area = func(width, length int) int {
	return width * length
}

func main()  {
    var width int = 24
	var length int = 34
	
	areaOfRectangle := area(width, length)
	fmt.Println(width, length, areaOfRectangle)
}
```

### Anonymous Function as Argument to another Function
```go
package main

import "fmt"

func main() {
	var width int = 12
	var length int = 23
	
	sum := func(w, l int) int {
		return w + l
    }
	
	multiply := func(w, l int) int {
		return w * l
    }
	
	perimeter := func(num int) int {
	    return num * 2	
    }
	
	area := func(num int) int {
		return num
    }
	
	
	rectangleArea := area(multiply(width, length))
	rectanglePerimeter := perimeter(sum(width, length))
	
	fmt.Printf("Area = %d and Perimeter = %d", rectangleArea, rectanglePerimeter)
}
```

### Anonymous Function as Return Value
```go
package main

import "fmt"

func displayNumber(number int) func() int  {
      return func() int {
		  number++
		return number
      }
}

func main()  {
  a := displayNumber(6)
  fmt.Println(a())
}
```
