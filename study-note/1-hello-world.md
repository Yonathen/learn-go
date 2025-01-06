# Writing Hello, World in Go
```go
package main

import "fmt"

func main()  {
    fmt.Println("Hello, World!")
}
```
- We organize go programs in packages
    - Declare which package it belongs
    - A package can be composed of multiple files or one
    - A program can contain multiple packages
    - `main` package is the entry point to the program
  - We use `import` keyword to import package
  - `fmt` is built-in package 
      - used for input/output
  - `main()` function is a special function
      - since its the place where the program starts
  - `fmt.Println("Hello world")`
      - formats accordint to the format specifier and writes to the standard output

- Compiling our go code 
  - `go run hello.go`
    - First compiles and then runs the program
  - `go build hello.go`
    - This will create a `hello` file
      - that's a binary you can execute