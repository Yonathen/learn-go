# Go Functions

Functions are a foundational concept in Go programming, enabling code reusability, modularity, and efficient organization. In Go, functions are blocks of statements grouped together to perform a specific task, which can be reused throughout the codebase.

---

## What is a Function?

- A **function** is a block of code that performs a specific operation.
- Functions are **not executed automatically**; they run only when **called**.
- Functions help in dividing large programs into smaller, manageable, and reusable pieces.

---

## Declaring a Function

To define a function in Go, use the `func` keyword as follows:

```go
func <function_name>(param1 <data_type1>, param2 <data_type2>) (<return_type1>, <return_type2>) {
    // Code block
    return value1, value2
}
```

- **function_name**: Name of the function.
- **params**: Input parameters with their data types.
- **return_type(s)**: Data type(s) of value(s) the function returns.
- A function may return multiple values in Go.

#### Example

```go
package main

import "fmt"

// calculate takes two integers and returns their sum and difference
func calculate(num1 int, num2 int) (int, int) {
    addition := num2 + num1
    subtraction := num2 - num1
    return addition, subtraction
}

func main() {
    var num1, num2 int

    fmt.Print("Enter num1: ")
    fmt.Scanf("%d", &num1)
    fmt.Print("Enter num2: ")
    fmt.Scanf("%d", &num2)

    sum, difference := calculate(num1, num2)
    fmt.Printf("The sum and difference of %d and %d are %d and %d respectively\n", num1, num2, sum, difference)
}
```

---

## Parameters and Arguments

- **Parameters** are variables listed in the function definition.
- **Arguments** are values supplied to the function when it is called.

Go supports:
- **Positional arguments** (order matters)
- **Named return values** (variables as return values)

#### Named Return Value Example

```go
func divide(a, b float64) (result float64, err error) {
    if b == 0 {
        err = fmt.Errorf("division by zero")
        return
    }
    result = a / b
    return
}
```

---

## Return Types

- Functions can return zero, one, or multiple values.
- If a function doesn't need to return anything, omit the return type.

#### Example: No return value
```go
func greet(name string) {
    fmt.Printf("Hello, %s!\n", name)
}
```

---

## Variable Scope

Variable scope defines the visibility and lifetime of a variable.

- **Local Variables**: Declared within a function, accessible only inside that function.
- **Global Variables**: Declared outside any function, visible throughout the package.

#### Example

```go
var globalVariable = "I am visible everywhere"

func someFunction() {
    localVariable := "I am visible only here"
    fmt.Println(globalVariable, localVariable)
}
```

---

## Function Types

Functions in Go can be categorized as:

1. **Named Functions**: Standard functions with names.
2. **Anonymous Functions**: Functions without names. Used as closures or for passing around functionality.

---

## Anonymous Functions (Lambdas / Closures)

An **anonymous function** is a function without a name. It is often assigned to a variable or passed as an argument.

#### Simple Anonymous Function

```go
package main

import "fmt"

var greet = func() {
    fmt.Println("Hi From Anonymous Func")
}

func main() {
    greet() // call the anonymous function
}
```

#### Anonymous Function with Parameters and Return Value

```go
var multiply = func(x, y int) int {
    return x * y
}

func main() {
    fmt.Println(multiply(10, 5)) // Output: 50
}
```

---

### Return Value from Anonymous Function

```go
var area = func(width, length int) int {
    return width * length
}

func main() {
    fmt.Println(area(6, 7)) // Output: 42
}
```

---

### Anonymous Function as Arguments

Anonymous functions can be passed as arguments to other functions:

```go
package main

import "fmt"

func processNumbers(a, b int, f func(int, int) int) int {
    return f(a, b)
}

func main() {
    add := func(x, y int) int {
        return x + y
    }
    fmt.Println(processNumbers(3, 7, add)) // Output: 10
}
```

---

### Anonymous Function as Return Value

Functions can return anonymous functions (this allows for closures):

```go
func nextNumber(start int) func() int {
    number := start
    return func() int {
        number++
        return number
    }
}

func main() {
    increment := nextNumber(5)
    fmt.Println(increment()) // Output: 6
    fmt.Println(increment()) // Output: 7
}
```

---

## Methods (Functions Attached to Types)

Go allows you to define functions as methods on user-defined types:

```go
type Rectangle struct {
    width, height int
}

// Area is a method with a receiver of type Rectangle
func (r Rectangle) Area() int {
    return r.width * r.height
}

func main() {
    rect := Rectangle{10, 5}
    fmt.Println(rect.Area()) // Output: 50
}
```

---

## Variadic Functions

Functions can accept a variable number of arguments:

```go
func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

func main() {
    fmt.Println(sum(1, 2, 3, 4, 5)) // Output: 15
}
```

---

## Defer, Panic, and Recover in Functions

- **defer**: Schedules a function call to run after the current function completes.
- **panic**: Terminates a function abruptly.
- **recover**: Regains control of a panicking goroutine.

#### Example

```go
func riskyOperation() {
    defer fmt.Println("Deferred execution")
    panic("Something went wrong!")
}

func main() {
    riskyOperation()
    fmt.Println("Will not execute due to panic")
}
```

---

## Key Points

- Functions promote code reusability and modularity.
- Go supports first-class functions: functions can be assigned to variables, passed as arguments, and returned from other functions.
- Functions can be named, anonymous, variadic, methods, or closures.

---

## Summary Table

| Function Type                   | Example Usage                                        |
|----------------------------------|-----------------------------------------------------|
| Named function                   | `func add(a, b int) int { return a + b }`           |
| Anonymous function               | `var f = func(a int) int { return a * 2 }`          |
| Method (type + receiver)         | `func (r Rectangle) Area() int { ... }`              |
| Variadic function                | `func sum(nums ...int) int { ... }`                  |
| Function as argument             | `func run(fn func()) { fn() }`                       |
| Function as return value         | `func gen() func() int { return func() int { ... }}` |

---

## Best Practices

- Start function names with a verb describing what they do.
- Keep functions focused on a single responsibility.
- Use return values for error handling (often with `(value, error)` pair).
- Clearly document parameter and return types if code is public.

---

## References

- [Go Official Documentation: Functions](https://go.dev/doc/effective_go#functions)
- [Go by Example: Functions](https://gobyexample.com/functions)
- [Effective Go](https://go.dev/doc/effective_go)
