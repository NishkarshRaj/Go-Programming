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
