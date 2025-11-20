# Go Data Types & Variables

Go (Golang) is a statically-typed, compiled language designed for simplicity, efficiency, and reliability. Understanding variables and data types is fundamental to writing idiomatic Go code.

---

## Basic Data Types

Go supports several built-in data types. Each type has specific characteristics and uses:

| Data Type | Description                      |
|-----------|----------------------------------|
| int       | Integer numbers                  |
| float     | Decimal numbers (floating point) |
| complex   | Complex numbers (e.g., 2+3i)     |
| string    | Sequence of characters           |
| bool      | true / false boolean values      |
| byte      | 8 bits of non-negative integer   |
| rune      | Unicode code point (int32)       |

> **Tip:** `byte` is an alias for `uint8`, commonly used to represent binary data; `rune` is an alias for `int32`, helpful for handling Unicode.

```go
var message string = "Welcome"
var boolValue bool = false
```

---

## Integer Data Types

Go provides both signed and unsigned integers in varying sizes:

| Data Type    | Size                        |
|--------------|-----------------------------|
| int/uint     | 32 bits (on 32-bit arch), 64 bits (on 64-bit arch) |
| int8/uint8   | 8 bits (1 byte)             |
| int16/uint16 | 16 bits (2 bytes)           |
| int32/uint32 | 32 bits (4 bytes)           |
| int64/uint64 | 64 bits (8 bytes)           |

```go
var integer int32 = 234
```

> **Note:** Use the smallest size necessary for your data to save memory.

---

## Floating-Point Data Types

Floating point numbers hold decimal values:

| Data Type | Size                |
|-----------|---------------------|
| float32   | 32 bits (4 bytes)   |
| float64   | 64 bits (8 bytes)   |

By default, Go treats floating-point literals as `float64`.

```go
var decimal float32 = 234.43
```

---

## Complex Numbers

Go has native support for complex numbers:

- `complex64`: float32 real and imaginary parts
- `complex128`: float64 real and imaginary parts

```go
var c complex64 = 1 + 2i
```

---

# Variables in Go

Variables are named memory locations used for storing data. Go variables must be declared before use.

## Types of Variables

- `int` : stores integers (whole numbers)
- `float32` : stores floating point numbers
- `string` : stores text
- `bool` : stores true or false

## Variable Declaration

- **Explicit declaration:**  
  `var <variable_name> <type> = <value>`
- **Type inference:**  
  `var <variable_name> = <value>` — Go infers type from value.
- **Short variable declaration:**  
  `<variable_name> := <value>` — only inside functions.

```go
package main

import "fmt"

func main() {
    var student1 string = "Jhon"
    var student2 = "Jane" // type inferred
    x := 2 // short declaration

    fmt.Println(student1)
    fmt.Println(student2)
    fmt.Println(x)
}
```

## `var` vs `:=`

| Feature      | `var`                          | `:=`                           |
|--------------|-------------------------------|--------------------------------|
| Where        | Inside & outside functions     | Only inside functions          |
| Separation   | Can declare without assignment | Must assign in declaration     |
| Readability  | Clear for package-level scope  | Concise for local scope        |

---

# Multiple Variable Declaration

You can declare multiple variables together for better readability and organization.

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

- **Inline declaration:**  
  `var a, b, c, d int = 1, 3, 5, 7` — all variables share the same type.
- **Mixed types without explicit type:**  
  `var e, f = 6, "Hello world"` & `g, h := 7, "Another World"`
- **Block declaration:**  
  Use for grouping variables conveniently.

---

# Variable Naming Rules

- Must start with a letter or underscore (_)
- Can contain letters, digits, underscores
- Case-sensitive
- No length limit
- Cannot be a Go keyword

### Naming conventions

- **Camel case:** `myVariableName := "Value"`
- **Pascal case:** `MyVariableName := "Value"`
- **Snake case:** `my_variable_name := "Value"`
- **Meaningful names:** Use clear, intention-revealing names.

> **Tip:** Use concise but descriptive names; for short-lived, temporary values, single letters like `i`, `j`, `x` are fine.

---

# Constants in Go

If you want a variable whose value cannot change, use `const`:

```go
const Pi = 3.14
const Greeting string = "Hello, Go!"
```

- Constants must be declared with a value up front.
- Type can be inferred or specified.

---

# Additional Go Insights

- **Zero Values:**  
  If a variable is declared but not assigned a value, it gets a "zero value":
  - Numbers: `0`
  - Strings: `""`
  - Booleans: `false`

- **Type Conversion:**  
  You can convert between compatible types:
  ```go
  var x int = 42
  var y float64 = float64(x)
  ```

- **Print Type:**  
  Use `fmt.Printf` with `%T` to print a variable's type:
  ```go
  fmt.Printf("Type of x is %T\n", x)
  ```

---

## Practice Task

Try declaring different types of variables and constants, print their values and types, and experiment with naming conventions.

```go
package main

import "fmt"

func main() {
    var age int = 30
    const country = "Ethiopia"
    score := 95.4
    fmt.Printf("Name: %s\n", country)
    fmt.Printf("Age: %d\n", age)
    fmt.Printf("Score: %.2f\n", score)
    fmt.Printf("Type of age: %T\n", age)
    fmt.Printf("Type of score: %T\n", score)
}
```

---

**Continue learning:**  
- Explore Go arrays, slices, maps, and structs to organize more complex data.
- Try writing functions with variables and see how scope works in Go.
- Look into Go tools like `go fmt` for auto-formatting your code.

Happy coding!
