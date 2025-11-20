# Go Constants

## What is a Constant?
- A constant is a *fixed, unchangeable value* that is determined at compile time.
- Once declared, the value **cannot be modified** throughout the program ("Unchangeable and readonly").
- Constants help ensure code reliability by preventing inadvertent modification of important, fixed values.

## General Rules
- Naming conventions for constants follow the same rules as variables (camelCase, PascalCase for exported names).
- Constants can be either **typed** or **untyped**.
  - *Typed constant*: Declared with an explicit type.
  - *Untyped constant*: Declared without specifying a type, may assume a different type depending on context.

## Declaration Syntax
Use the `const` keyword, either outside or inside a function:

```go
const typedConst int = 100
const untypedConst = "Hello"
```

Constants can also be grouped:

```go
const (
    Pi        = 3.14159
    Language  = "Go"
    DaysInWeek = 7
)
```

## Example Program

```go
package main

import "fmt"

// Package-level constants
const TYPED_CONSTANT int = 1
const UNTYPED_CONSTANT = 2

func main() {
    // Block constants within a function
    const (
        A   int = 1
        PIE      = 3.14
        C        = "Area of a circle"
    )

    fmt.Println(TYPED_CONSTANT, UNTYPED_CONSTANT)
    fmt.Println(A, PIE, C)
}
```

## Key Properties & Missing Details

### Allowed Types
- Constants can only be of:
  - Boolean (`true`, `false`)
  - Numeric types (integer, floating-point, complex)
  - String
- **Slices, maps, structs, arrays, and functions** cannot be constant.

### Implicitly Untyped Constants
- Untyped constants can be used as any compatible type in assignment.
  ```go
  const x = 42     // Untyped
  var y int = x    // Type inferred as int
  var z float64 = x // Type inferred as float64
  ```
- Makes them flexible and useful in expressions.

### Compile-Time Evaluation
- Constants must be assigned with compile-time constant expressions only.
- You cannot assign runtime values or expressions to a constant:
  ```go
  const myConst = math.Sqrt(4) // ❌ Invalid, math.Sqrt runs at runtime
  ```

### Enumerated Constants and Iota
- Go uses `iota` for incrementing values, typically for enums.
- `iota` starts at zero in each const block and increments automatically.
- Example:

  ```go
  const (
      Sunday = iota
      Monday
      Tuesday
      Wednesday
      Thursday
      Friday
      Saturday
  )
  fmt.Println(Sunday, Monday, Saturday) // Output: 0 1 6
  ```

### Exporting Constants
- Constants starting with an uppercase letter are **exported** (visible outside the package).
  ```go
  const MaxLength = 1024      // Exported Constant
  const minLength = 1         // Unexported (internal) Constant
  ```

### Constant Expressions
- You can use arithmetic and logical operators to form constant expressions:
  ```go
  const radius = 5
  const area = Pi * radius * radius
  ```

## Best Practices

- Prefer constants for values that never change and are referenced in multiple places.
  - Example: Mathematical values, limits, configuration settings, keys, enum values.
- Use `iota` for enumerated sequences to avoid manual value assignment.
- Give constants descriptive names that clearly indicate their purpose.

---

## Summary Table

| Aspect         | Details                                               |
| -------------- | ---------------------------------------------------- |
| Mutability     | Immutable (cannot change after declaration)          |
| Types allowed  | bool, numeric, string                                |
| Scope          | Package-level or block-level                         |
| Exported       | Name starts with uppercase letter                    |
| Special tools  | `iota` for sequences/enums                           |
| Runtime usage  | Compile-time only                                    |

## Reference

- [Go constants - Official Tour](https://tour.golang.org/basics/15)
- [Effective Go: Constants](https://go.dev/doc/effective_go#constants)
