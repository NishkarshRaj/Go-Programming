# Decision Making in Golang

1. If

```go
package main
import "fmt"
func main() {
if (41 > 0 ) {
fmt.Println("Inside If");
}
}
```

2. Else if

```go
package main
import "fmt"
func main() {
a := 10 // dynamic declaration
if(a<10) {
fmt.Println("Less than 10")
} else if (a>10) {
fmt.Println("Greater than 10")
} else {
fmt.Println("Equal to 10")
}

// else if and else must come in same line due to EOL strict mapping of Go
}
```

3. Nested if

```go
package main
import "fmt"
func main() {
a :=41
if(a>10) {
if(a<100) {
fmt.Println("In Range 10-100")
} else {
fmt.Println("Greater than or equal to 100")
}
} else {
fmt.Println("Less than 10")
}
}
```


4. Switch Statement (NOT A FUNCTION)

4.1 Expressional Switch

* Expressions must result in boolean expression or integer type - default is true
* No break is needed unlike C, if matched, only that case is executed 
* 

```go
package main
import "fmt"

func main() {
   /* local variable definition */
   var grade string = "B"
   var marks int = 90
  
    // switch wrt marks
   switch marks {
      case 90: grade = "A"
      case 80: grade = "B"
      case 50,60,70 : grade = "C" // multiple comparisions allowed
      default: grade = "D"  
   }
   switch { // switch wrt true - must compare within case
      case grade == "A" :
         fmt.Printf("Excellent!\n" )     
      case grade == "B", grade == "C" :
         fmt.Printf("Well done\n" )      
      case grade == "D" :
         fmt.Printf("You passed\n" )      
      case grade == "F":
         fmt.Printf("Better try again\n" )
      default:
         fmt.Printf("Invalid grade\n" );
   }
   fmt.Printf("Your grade is  %s\n", grade );      
}
```

4.2 Type Switch

* Condition statement must return interface{} type x.(type)

```go
package main
import "fmt"

func main() {
   // declare interface x
   var x interface{}
     
   switch i := x.(type) { // dynamic declaration of variable I of interface type
      case nil:	  
         fmt.Printf("type of x :%T",i)                
      case int:	  
         fmt.Printf("x is int")                       
      case float64:
         fmt.Printf("x is float64")           
      case func(int):
         fmt.Printf("x is func(int)")                      
      case bool, string:
         fmt.Printf("x is bool or string")       
      default:
         fmt.Printf("don't know the type")     
   }   
}
```


5. Select Statement

**Syntax:**

```
select {
   case communication clause  :
      statement(s);      
   case communication clause  :
      statement(s); 
   /* you can have any number of case statements */
   default : /* Optional */
      statement(s);
}
```

* No expression after select statement
* COMMUNICATION CHANNEL CLAUSE FOLLOWS "CHAN" keyword of Golang

```go
package main

import "fmt"

func main() {
   var c1, c2, c3 chan int
   var i1, i2 int
   select {
      case i1 = <-c1:
         fmt.Printf("received ", i1, " from c1\n")
      case c2 <- i2:
         fmt.Printf("sent ", i2, " to c2\n")
      case i3, ok := (<-c3):  // same as: i3, ok := <-c3
         if ok {
            fmt.Printf("received ", i3, " from c3\n")
         } else {
            fmt.Printf("c3 is closed\n")
         }
      default:
         fmt.Printf("no communication\n")
   }    
}   
```
