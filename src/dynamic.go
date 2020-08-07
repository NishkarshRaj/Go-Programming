package main
import "fmt"

func main() {
var x uint;
x = 41; //static requires var declaration
y := 31; //dynamic - dont specify data type - let compiler decide optimally - runtime
fmt.Println(x,y)
}
