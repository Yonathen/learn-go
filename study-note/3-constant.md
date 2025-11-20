# Go Input & Output Functions

Go provides powerful facilities for taking input from users and displaying output. Understanding these functions is essential for effective interaction in CLI applications.

---

## Input Functions in Go (`fmt` Package)

Go offers multiple variations of the `scan` methods for reading user input from the standard input, each with its own use cases.

### Variations of Scan Functions

- [`fmt.Scan()`](https://pkg.go.dev/fmt#Scan)
- [`fmt.Scanln()`](https://pkg.go.dev/fmt#Scanln)
- [`fmt.Scanf()`](https://pkg.go.dev/fmt#Scanf)

---

### 1. `fmt.Scan()`

- Reads input separated by spaces.
- Stops reading variables at whitespace.
- Suitable for quickly reading multiple space-separated values.

**Example:** *(scan name and age separated by a space)*
```go
package main

import "fmt"

func scanNameAndAge() {
	fmt.Println("\nScan Name And Age By scan()")
	var name string
	var age int

	fmt.Print("Enter your name and age: ")
	fmt.Scan(&name, &age)

	fmt.Printf("Hi %s, Your age is %d\n", name, age)
}

func main() {
	scanNameAndAge()
}
```

**Sample Input:** `Alice 29`

**Sample Output:** `Hi Alice, Your age is 29`

---

### 2. `fmt.Scanln()`

- Reads input until a newline character is encountered.
- Useful for taking multiple values, but stops reading when Enter is pressed.

**Example:**
```go
package main

import "fmt"

func scanlnNameAndAge() {
	fmt.Println("\nScan Name And Age By scanln()")
	var name string
	var age int
	fmt.Print("Enter your name and age: ")
	fmt.Scanln(&name, &age)
	fmt.Printf("Hi again %s, this is scanln. Your age is %d\n", name, age)
}

func main() {
	scanlnNameAndAge()
}
```

**Sample Input:** `Bob 35`

---

### 3. `fmt.Scanf()`

- Reads formatted input using format specifiers (just like `printf` for output).
- Flexible for customized parsing.

**Example:**
```go
package main

import "fmt"

func scanfNameAndAge() {
	fmt.Println("\nScan Name And Age By scanf()")
	var name string
	var age int
	fmt.Print("Enter your name and age: ")
	fmt.Scanf("%s %d", &name, &age)
	fmt.Printf("Hi again %s, this is scanf. Your age is %d\n", name, age)
}

func main() {
	scanfNameAndAge()
}
```

**Sample Input:** `Charlie 44`

---

### Notes

- All these scan functions return the number of items successfully scanned and an error (if any).
- Always check for errors in production programs for robust input handling.

**Example:**
```go
n, err := fmt.Scan(&name, &age)
if err != nil {
    fmt.Println("Error reading input:", err)
}
```

---

## Output Functions in Go (`fmt` Package)

Go provides several functions for formatted output. The main ones are:

- `fmt.Print()`
- `fmt.Println()`
- `fmt.Printf()`

---

### 1. `fmt.Print()`

- Prints arguments as-is, no extra spaces added, and doesn’t append a newline.
- To print on a new line, explicitly add `\n`.

**Example:**
```go
fmt.Print("Hello")
fmt.Print("World")
fmt.Print("\n") // For newline
```
**Output:** `HelloWorld`

---

### 2. `fmt.Println()`

- Prints arguments separated by a space and automatically appends a newline.

**Example:**
```go
fmt.Println("Hello")
fmt.Println("World")
```

**Output:**
```
Hello
World
```

---

### 3. `fmt.Printf()`

- Prints formatted output according to the specified format string and formatting verbs.

**Example:**
```go
var name = "David"
var age = 25
fmt.Printf("Hi, %s! You are %d years old.\n", name, age)
```

---

## Output Example: All Three Functions

```go
package main

import "fmt"

func main() {
	var str1, str2 = "Print : Prints", "On Same lines"

	fmt.Print(str1)
	fmt.Print(str2)
	fmt.Print("\n", str1, "\n", str2, "\n")

	var str3, str4 = "Println : Prints", "Always on the new line"
	fmt.Println(str3)
	fmt.Println(str4)

	var (
		str5        = "i has value = %v and type %T\n"
		str6        = "j has value = %v and type %T\n"
		i    string = "Hello Again"
		j    int    = 20
	)
	fmt.Printf(str5, i, i)
	fmt.Printf(str6, j, j)
}
```

---

## Formatting Verbs Reference

The `fmt.Printf()` and `fmt.Scanf()` functions use *verbs* for formatting. Here are the most commonly used ones:

### General Formatting
| Verb   | Description                        |
|--------|------------------------------------|
| `%v`   | Value in default format            |
| `%#v`  | Value in Go-syntax                 |
| `%T`   | Type of the value                  |
| `%%`   | Literal percent sign               |

### Strings
| Verb    | Description                           |
|---------|---------------------------------------|
| `%s`    | Plain string                          |
| `%q`    | Double-quoted string                  |
| `%8s`   | Width 8, right justified              |
| `%-8s`  | Width 8, left justified               |
| `%x`    | Hexadecimal                          |
| `% X`   | Hexadecimal with spaces               |

### Integers
| Verb    | Description                                        |
|---------|----------------------------------------------------|
| `%b`    | Base 2 (binary)                                    |
| `%d`    | Base 10 (decimal)                                  |
| `%+d`   | Always show sign                                   |
| `%o`    | Base 8 (octal)                                     |
| `%O`    | Base 8 (with leading 0o)                           |
| `%x`    | Base 16, lowercase                                 |
| `%X`    | Base 16, uppercase                                 |
| `%#x`   | Base 16, with leading 0x                           |
| `%4d`   | Right justified, width 4                           |
| `%-4d`  | Left justified, width 4                            |
| `%04d`  | Pad with zeroes, width 4                           |

### Floats
| Verb     | Description                                         |
|----------|-----------------------------------------------------|
| `%e`     | Scientific notation (e.g., -1.234e+06)              |
| `%f`     | Decimal format (default precision)                  |
| `%.2f`   | Two digits after the decimal point                  |
| `%6.2f`  | Width 6, precision 2                                |
| `%g`     | Compact (exponent only if needed)                   |

---

## Input/Output Best Practices

- **Always handle errors!** For example, reads might fail or get incomplete input.
- Use **comments** and clear prompts to help users.
- For repeated input, consider using **loops** and **bufio.Scanner** for more control and flexibility.

---

## Advanced: Buffered Input With `bufio`

For line-by-line input or better performance, use the `bufio` package.

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter a line: ")
    input, _ := reader.ReadString('\n')
    fmt.Println("You entered:", input)
}
```

---

## Summary Table

| Function    | Reads until | Format Support | Example Use Case      |
|-------------|-------------|----------------|-----------------------|
| `Scan()`    | Space       | No             | Small set of words    |
| `Scanln()`  | Newline     | No             | Multi-word line entry |
| `Scanf()`   | Custom fmt  | Yes            | Structured input      |
| `Print()`   | -           | No             | Simple output         |
| `Println()` | -           | No             | Output with newline   |
| `Printf()`  | -           | Yes            | Formatted output      |

---

### Additional Resources

- [Official fmt Documentation](https://pkg.go.dev/fmt)
- [Go by Example: Input](https://gobyexample.com/reading-files)
- [Tour of Go: Formatting](https://tour.golang.org/basics/15)

---

> **Tip:** For more comprehensive input handling, explore `bufio.Scanner` and handling errors gracefully.
