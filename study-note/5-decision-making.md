# Decision-Making in Go

Go provides robust decision-making capabilities that enable developers to control the flow of application logic. This document explores the main constructs available for decision-making in Go, including `if`, `switch` statements and the various looping mechanisms. Each section includes syntax details, best practices, and code examples to deepen your understanding.

---

## 🟩 If Statements

The `if` statement is the primary tool for conditional logic in Go. It evaluates expressions and executes code blocks based on the truthiness of those expressions.

### **Supported Operators**

- **Comparison Operators:** `<`, `>`, `<=`, `>=`, `==`, `!=`
- **Logical Operators:** `&&` (AND), `||` (OR), `!` (NOT)

### **Syntax**

```go
if condition_1 {
    // Executes if condition_1 is true
} else if condition_2 {
    // Executes if condition_2 is true
} else {
    // Executes if none of the above conditions are true
}
```

### **Example**

```go
package main

import "fmt"

func greetByTime(time int) {
    fmt.Println("\nIf Condition Example")
    if time < 10 {
        fmt.Println("Good morning")
    } else if time > 12 {
        fmt.Println("Good Afternoon")
    } else {
        fmt.Println("Good evening")
    }
}

func main() {
    greetByTime(3)   // Output: Good morning
    greetByTime(13)  // Output: Good Afternoon
}
```
- **Tip:** Use `else if` chains for multi-way decisions.

---

## 🟨 Switch Statement

Go’s `switch` statement is a concise way to compare an expression against multiple cases. Unlike other languages, Go’s `switch` does **not** require *break* statements as cases do not fall through by default.

### **Syntax**

```go
switch expression {
    case x:
        // Executes if expression == x
    case y, z:
        // Executes if expression == y or z
    default:
        // Executes if no cases match
}
```

### **Example**

```go
package main

import "fmt"

func dayType(day string) {
    switch day {
    case "Saturday", "Sunday":
        fmt.Println("It's Weekend!")
    case "Monday":
        fmt.Println("It's Monday, back to work.")
    default:
        fmt.Println("It's a week day.")
    }
}

func main() {
    dayType("Saturday")
    dayType("Wednesday")
}
```

- **Note:** You can list multiple values in a single case.

---

## 🔁 Looping Constructs

Go uses the `for` keyword for all looping operations. The `for` loop can imitate the behavior of traditional `for`, `while`, or `do...while` loops from other languages.

### 1. While-Style Loop

Executes a code block **while** the given condition is `true`.

#### **Syntax**

```go
for condition {
    // Code Block
}
```

#### **Example**

```go
package main

import "fmt"

func main() {
    i := 0
    for i < 10 {
        fmt.Printf("i = %d\n", i)
        i++
    }
}
```

---

### 2. Simulating do...while Loop

Go doesn’t have a built-in `do...while` structure, but you can use an infinite `for` loop with a logical `break` statement:

#### **Example**

```go
package main

import "fmt"

func main() {
    i := 0
    for {
        if i > 10 {
            break
        }
        fmt.Printf("i = %d\n", i)
        i++
    }
}
```

---

### 3. Classical For Loop

This is the most common type of loop, initializing a value, checking a condition, and incrementing the counter at every iteration.

#### **Syntax**

```go
for initialization; condition; increment {
    // Code block
}
```
- **initialization**: executes once when the loop starts.
- **condition**: checked before every iteration.
- **increment**: executes at the end of each loop iteration.

#### **Range-Based For Loop**

Ideal for iterating over arrays, slices, maps, strings, and channels:

```go
for index, value := range collection {
    // Code block
}
```
- **Tip:** Use `_` if you do not need the index or value.

#### **Example**

```go
package main

import "fmt"

func main() {
    fmt.Println("For loop: continue and break demo")
    for i := 0; i <= 10; i++ {
        if i == 8 {
            break // Exit loop early
        }
        if i == 3 {
            continue // Skip to next iteration
        }
        fmt.Printf("Binary: %b\n", i)
        fmt.Printf("Hex: %x\n", i)
        fmt.Printf("Decimal: %d\n", i)
    }

    fmt.Println("\nNested loop example:")
    adjectives := [2]string{"Big", "Tasty"}
    fruits := [3]string{"Apple", "Orange", "Banana"}
    for _, adj := range adjectives {
        for _, fruit := range fruits {
            fmt.Printf("Kind: %s, Fruit: %s\n", adj, fruit)
        }
    }

    fmt.Println("\nIterating with Range (show index):")
    for index, value := range fruits {
        fmt.Printf("Value: %s at index %d\n", value, index)
    }

    fmt.Println("\nIterating with Range (hide index):")
    for _, value := range fruits {
        fmt.Printf("Value: %s\n", value)
    }
}
```

---

## 💡 Good Practices

- Prefer `switch` when comparing a variable against known values.
- Use `continue` and `break` to control loop execution.
- Use range-loop for collection-based iteration.
- Keep conditions readable and simple.

---

## 🎓 Summary

Go streamlines decision-making with its clear and concise syntax for conditional and loop constructs. Mastering these will help you write expressive, efficient, and robust Go programs. Experiment with nested and range-based loops, tailor logic branches with `if` and `switch`, and you'll be able to handle any application flow with confidence!
