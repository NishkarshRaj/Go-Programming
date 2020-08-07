# Operators in Golang

1. Arithmetic

2. Relational

3. Logical

4. Bitwise

5. Assignment

6. Miscellaneous Operators

## Code

```go
package main
import "fmt"
func main(){

// variable declaration
var a,b = 10,20

// 1. Arithmetic operators
fmt.Printf("a+b: %d\n",a+b)
fmt.Printf("a-b: %d\n",a-b)
fmt.Printf("a*b: %d\n",a*b)
fmt.Printf("a/b: %d\n",a/b)
fmt.Printf("a mod b: %d\n",a%b)
// fmt.Printf("a++: %d\n",a++) // not working
// fmt.Printf("a--: %d\n",a--)

// 2. Logical Operators

// 3. Relational Operators 

// 4. Bitwise Operators
fmt.Printf("a&b: %d\n",a&b) // bitwise and
fmt.Printf("a|b: %d\n",a|b) // bitwise or
fmt.Printf("a^b: %d\n",a^b) // bitwise xor
fmt.Printf("a<<2: %d\n",a<<2) // left shift mul 2
fmt.Printf("a>>2: %d\n",a>>2) // right shift div 2

// 5. Assignment Operators

// 6. Miscellaneous operators - sizeof and :?
// Pointers
var ptr *int // pointer declaration
ptr = &a // ptr assignment
fmt.Printf("Address of A: %d\n",ptr)
fmt.Printf("Dereferenced value of A: %d\n",*ptr)
}
```
