# Go Input Functions
- Go has 3 different variation of `scan()` method
  - `fmt.Scan()`
  - `fmt.Scanln()`
  - `fmt.Scanf()`

## Go fmt.Scan()
- Takes input from user
- It can only take input values up to a space
  - So when it counters a space the value before the space is assigned

- Example: To take multiple input from the user look at the following example
```go
package main

import "fmt"

func scanNameAndAge() {
	fmt.Println("\nScan Name And Age By scan()")
	var name string
	var age int

	fmt.Print("Enter your name and age ")
	fmt.Scan(&name, &age)

	fmt.Printf("\n Hi %s, Your age is %d \n", name, age)

}
func main() {
	scanNameAndAge()
}
```

## Go fmt.Scanln()
- get the input values up to the new line
- When it encounters a new line it stops taking input
```go
package main

import "fmt"

func scanlnNameAndAge() {
	fmt.Println("\nScan Name And Age By scanln()")
  var name string
  var age int
    fmt.Print("Enter your name and age ")
    fmt.Scanln(&name, &age)
    fmt.Printf("\n Hi again %s, this is scanln. Your age is %d \n", name, age)
}
func main() {
    scanlnNameAndAge()
}
```
## Go fmt.scanf()
- Similar to `scanln()`
  - Except in this case it takes inputs using format specifier
```go
package main

import "fmt"

func scanlnNameAndAge() {
	fmt.Println("\nScan Name And Age By scanln()")
  var name string
  var age int
    fmt.Print("Enter your name and age ")
    fmt.Scanf("%s %d", &name, &age)
    fmt.Printf("\n Hi again %s, this is scanf. Your age is %d \n", name, age)
}
func main() {
    scanlnNameAndAge()
}
```

# GO Output Functions
- Has three functions
  - `Print()`
  - `Println()`
  - `Printf()`

- Consider the following
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

## Print() function
- prints its arguments with their default format on the same line
- use `\n` as an argument if you want new line between the texts
- There is no whitespace between the argument values

## Println() function
- Similar to `Print()` except `Println`
  - Adds whitespace between argument values
  - Adds newline at the end

## Printf() function
- First formats the arguments based on the formatting verb
- Then prints the result

## General Formatting Verbs
- The following verbs can be used with all data types:
    - `%v`	Prints the value in the default format
    - `%#v`	Prints the value in Go-syntax format
    - `%T`	Prints the type of the value
    - `%%`	Prints the % sign

### String Formatting Verbs
- The following verbs can be used with the string data type:
    - `%s`	Prints the value as plain string
    - `%q`	Prints the value as a double-quoted string
    - `%8s`	Prints the value as plain string (width 8, right justified)
    - `%-8s`	Prints the value as plain string (width 8, left justified)
    - `%x`	Prints the value as hex dump of byte values
    - `%x`	Prints the value as hex dump with spaces

## Integer Formatting Verbs
- The following verbs can be used with the integer data type:
  - `%b`	Base 2
  - `%d`	Base 10
  - `%+d`	Base 10 and always show sign
  - `%o`	Base 8
  - `%O`	Base 8, with leading 0o
  - `%x`	Base 16, lowercase
  - `%X`	Base 16, uppercase
  - `%#x`	Base 16, with leading 0x
  - `%4d`	Pad with spaces (width 4, right justified)
  - `%-4d`	Pad with spaces (width 4, left justified)
  - `%04d`	Pad with zeroes (width 4)

## Float Formatting Verbs
- The following verbs can be used with the float data type:
    - `%e`	Scientific notation with 'e' as exponent
    - `%f`	Decimal point, no exponent
    - `%.2f`	Default width, precision 2
    - `%6.2f`	Width 6, precision 2
    - `%g`	Exponent as needed, only necessary digits