# Looping in Golang

## For loops

```go
package main
import "fmt"

func main() {
   var b int = 15
   var a int
   numbers := [6]int{1, 2, 3, 5}  // dynamic declaration of an array

   // type 1: C type definition
   for a := 0; a < 10; a++ {
      fmt.Printf("value of a: %d\n", a)
   }
   // Note since a value was dynamically assigned within loop, outside loop a = 0
   // if a = 0 would have been used a = 10 would have been the current value
   
   // type 2: conditional loop - runs till true - empty condition means true - infinite loop
   for a < b {
      a++
      fmt.Printf("value of a: %d\n", a)
   }
   
   // type 3: Ranged loop
   // i is automatically incremented from 0
   for i,x:= range numbers {
      fmt.Printf("value of x = %d at %d\n", x,i)
   }   
}
```

## Nested loops

```go
package main
import "fmt"
func main() {
var i,j int;
for i=1;i<=4;i++ {
for j=1;j<=i;j++ {
fmt.Printf("*")
}
fmt.Printf("\n")
}
}
```

## Loop Control Statements

### 1. Continue

### 2. Break

### 3. Goto

## Infinite Loops

```go
package main
import "fmt"
func main() {
// for { or
// for true {
for ; ; {
fmt.Println("Infinite Loop!");
}
}
```
