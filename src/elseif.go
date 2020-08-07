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
