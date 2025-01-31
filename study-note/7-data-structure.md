# Chapter 7 : Data Structures

## Go Arrays
- Is a collection of similar types of data
  - Syntax
    `var <array_variable> = [<size>]datatype{<elements_of_array>}`
    - Example
    ```go
        package main
    
        import "fmt"
    
        func main() {
            var arrayOfInteger = [5]int{1, 2, 3, 4, 5}
            
            fmt.Println(arrayOfInteger)
        }
    ```
    
### Accessing array elements in Golang
- Each element is associated to a number
    ```go
        package main
        
        import "fmt"
        
        func main() {
            integers := [5]int{1, 2, 3, 4, 5}
            
            fmt.Println(integers[0])
            fmt.Println(integers[1])
        }
    ```
  
### Initializing array elements
- In Go we can initialize specific element of the array during declaration
    ```go
        package main
        
        import "fmt"
        
        func main() {
            integers := [5]int{0:1, 3:2, 4:3, 2:4, 1:5}
            var integers1 [3]int
        
            integers1[0] = 1
            integers1[1] = 2
            integers1[2] = 3
  
            fmt.Println("Init specific element", integers)
            fmt.Println("Normal Init", integers1)
        
        }
    ```
  
### Changing the array element in go
```go
package main

import "fmt"

func main() {
	integers2 := [...]int{1, 3, 7, 9, 12}

	integers2[2] = 8
	
	fmt.Println(len(integers2), integers2[2])
}

```

### Looping through an array and multidimensional array
```go
package main

import "fmt"

func main() {
	integers2 := [5]int{1, 3, 7, 9, 12}
	integers2D := [2][2]int{{1, 2}, {2, 3}}


	for i := 0; i < len(integers2); i++ {
		fmt.Println(integers2[i])
	}
	
	for index, value := range integers2 {
		fmt.Println(value, index)
    }
	
	i := 0
	for i < len(integers2) {
		fmt.Println(integers2[i])
		i++
    }
	
	for _, value := range integers2 {
		fmt.Print(value)
	}

	fmt.Println("\n")
	for _, value := range integers2D {
		for _, item := range value {
			fmt.Print(item)
		}
	}
}

```

## Go Slice
- Same as array
  - except in this case it does not have fixed size
- Syntax
  `[]<type>{<elements>}`

