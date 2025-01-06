# Go Variables
- Variables are containers for storing data values

## Go Variable types
- Some of the types are
  - `int` : stores integers (whole numbers)
  - `float32` : stores floating point numbers
  - `string` : stores texts
  - `bool` : Stores true or false

## Variable Declarations
- Syntax
  - `var <variable_name> <type> = <value>`
- If you use `:=` sign as follows it means that the type of the variable is inferred from a value
  - `<variable_name> := value`

```go
package main

import "fmt"

func main() {
    var student1 string = "Jhon"
    var student2 = "Jane"
    x := 2

    fmt.Println(student1)
    fmt.Println(student2)
    fmt.Println(x)
}
```

## `var` Vs `:=`
- `var` can be used inside and outside functions
  - `:=` Can only be used inside functions
- In `var` Declaration and value assignment can be done separately
  - Whereas `:=` Must be done in the same line

# Multiple Variable Declaration
```go
package main

import "fmt"

func main() {

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
```
- `var a, b, c, d int = 1, 3, 5, 7`
  - Shows how to declare multiple variables in the same line
  - Here we have used a type keyword `int`
    - Therefor the only thing we can assign is integer
- `var e, f = 6, "Hello world"` & ` g, h := 7, "Another World"`
  - Here, in both case type is not specified
    - Therefore, you can declare different kind of variables in the same line
- The following is a Go Variable Declaration in a block
```go
  var (
    i int
    j int    = 1
    k string = "hello"
  )
```
  - For greater readability multiple variable declarations can be grouped together into a block

# Variable naming rules
- must start with letters or an underscore character (_)
- Can only contain alphanumeric characters and underscore
- they are case-sensitive
- no limit on length
- can not be any of the Go keywords

- Examples
  - Camel case : `myVariabName := "Value"`
  - Pascal case : `MyVariableName := "Value"`
  - Snack case : `My_Variable_Name := "Value"`

