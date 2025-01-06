# Go Constants
- Should have fixed value that cannot be changed
- "Unchangeable and readonly"
- Follows the same naming rule as variables
- Consider the following code
```go
package main

import "fmt"

const TYPED_CONSTANT int = 1
const UNTYPED_CONSTANT = 2

func main() {
	const (
		A   int = 1
		PIE     = 3.14
		C       = "Area of a circle"
	)

	fmt.Println(TYPED_CONSTANT, UNTYPED_CONSTANT)
	fmt.Println(A, PIE, C)
}
```