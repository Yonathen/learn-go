# Chapter 7 : Advanced Data Structures in Go

---

## 1. Arrays in Go

### What is an Array?
An **array** is a fixed-length, ordered collection of elements of the same type. Arrays in Go have the following characteristics:
- The length is part of the type, e.g., `[5]int` and `[10]int` are different types.
- Arrays cannot be resized once declared.
- Memory allocation is contiguous.

#### Declaration and Initialization
In Go, you can declare and initialize arrays in several ways.

**Syntax:**
```go
var <arrayName> = [<size>]<Type>{<elements>}
```

**Example:**
```go
package main

import "fmt"

func main() {
    var numbers = [5]int{1, 2, 3, 4, 5}
    fmt.Println(numbers)
}
```

### Accessing Array Elements
Each element can be accessed with its **zero-based index**.

```go
package main

import "fmt"

func main() {
    arr := [5]int{1, 2, 3, 4, 5}
    fmt.Println(arr[0]) // 1
    fmt.Println(arr[4]) // 5
}
```

### Partial Initialization & Assignment
You can initialize specific elements by index:

```go
package main

import "fmt"

func main() {
    arr := [5]int{0: 10, 3: 30} // [10 0 0 30 0]
    arr[1] = 20                 // [10 20 0 30 0]
    fmt.Println(arr)
}
```

#### Declaring Without Initialization
```go
var arr [3]int // Initialized with zeros: [0 0 0]
arr[0] = 7
arr[1] = 12
arr[2] = 28
```

### Modifying Elements
```go
package main

import "fmt"

func main() {
    arr := [...]int{1, 3, 7, 9, 12} // Length auto-detected
    arr[2] = 8
    fmt.Printf("Length: %d, Changed Value: %d\n", len(arr), arr[2])
}
```

---

### Iterating Over Arrays

#### Traditional For Loop
```go
for i := 0; i < len(arr); i++ {
    fmt.Println(arr[i])
}
```

#### Using `range`
```go
for idx, val := range arr {
    fmt.Printf("Index: %d, Value: %d\n", idx, val)
}
```

#### While-like Loop
```go
i := 0
for i < len(arr) {
    fmt.Println(arr[i])
    i++
}
```

#### Ignoring Index
```go
for _, val := range arr {
    fmt.Println(val)
}
```

#### Multidimensional Arrays
Go supports multi-dimensional arrays:
```go
matrix := [2][2]int{{1, 2}, {3, 4}}
for _, row := range matrix {
    for _, cell := range row {
        fmt.Print(cell, " ")
    }
}
fmt.Println()
```

---

## 2. Slice: Dynamic Data Structure

### What is a Slice?
A **slice** is a dynamically-sized, flexible view into an array. It is a reference type, and changes to a slice affect the underlying array.

- Slices can grow or shrink.
- Slices have a length and a capacity.
- The zero value of a slice is `nil`.

**Syntax of a Slice:**
```go
var name []Type
```

**Example:**
```go
numbers := []int{1, 2, 3, 4, 5}
fmt.Println(numbers, len(numbers), cap(numbers))
```

---

### Advanced Array and Slice Operations

#### Appending Elements to a Slice
```go
s := []int{}
s = append(s, 1, 2, 3)
```

#### Slicing Arrays and Slices
```go
arr := [5]int{1, 2, 3, 4, 5}
sub := arr[1:4] // Elements at index 1,2,3: [2 3 4]
```

#### Copying Slices
```go
source := []int{1, 2, 3}
dest := make([]int, len(source))
copy(dest, source)
```

#### Multi-dimensional Slices
```go
matrix := [][]int{{1, 2}, {3, 4}, {5, 6}}
for i, row := range matrix {
    fmt.Printf("Row %d: %v\n", i, row)
}
```

#### Capacity and Length
```go
a := make([]int, 5)      // len(a)=5, cap(a)=5
b := make([]int, 0, 10)  // len(b)=0, cap(b)=10
```

#### Removing Elements (Idiomatic Way)
```go
a := []int{1, 2, 3, 4, 5}
a = append(a[:2], a[3:]...) // Remove a[2], result: [1 2 4 5]
```

---

## Comparison: Array vs Slice

| Feature      | Array                   | Slice                       |
| ------------ | ----------------------- | -------------------------- |
| Length       | Fixed                   | Dynamic                    |
| Memory       | Value type              | Reference type             |
| Zero value   | All elements zero       | nil                        |
| Declaration  | `[N]Type`               | `[]Type`                   |
| Resizing     | Not possible            | Possible via `append`      |
| Useful For   | Performance, constants  | Flexibility, most cases    |

---

## Further Reading

- [Go Official Tour: Arrays](https://go.dev/tour/moretypes/6)
- [Effective Go: Slices](https://go.dev/doc/effective_go#slices)
- [Go by Example: Arrays and Slices](https://gobyexample.com/arrays)

---

> **Tip:** In real-world Go code, slices are favored over arrays due to their dynamic nature and ease of use. Arrays occasionally appear for fixed-size buffers or explicit memory layouts.
