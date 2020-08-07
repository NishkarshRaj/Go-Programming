package main
import "fmt"

func main() {
var x uint8 = 41 //specify data type to compiler - compile time
fmt.Println(x," Variable value: ",x) //concatenation of values and strings
fmt.Printf("Variable is of type %T\n",x) // Type - %T
}
