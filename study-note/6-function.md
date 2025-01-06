# Go Functions

- is a block of statements that can be used repeatedly
- will not execute automatically
- executed by a call to the function

## Creating a function
- Use a `func` keyword
- Syntax
  ```
  func <function_name>() {
    // Code block
  }
  ```
  
- Example
```go
package main
import ("fmt")

func myMessage() {
  fmt.Println("I just got executed!")
}

func main() {
  myMessage()
  myMessage()
  myMessage()
}
```