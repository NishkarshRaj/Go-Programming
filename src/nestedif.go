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
