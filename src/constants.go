package main
import "fmt"
func main(){

// 1. Integer constants
fmt.Println(41) // integer constant [0-9]ul // unsigned and long can be capital or small and can appear in any order at suffix
fmt.Println(034) // octal constant 0[0-7]ul
fmt.Println(0xF2) // Hexadecimal constant 0[x/X](0-9,A-F)ul

// 2. Floating constants
fmt.Println(41.31) // decimal constants
fmt.Println(2e5) // exponential constant

// 3. String Literals
fmt.Println("Hello, World!")
fmt.Println("Hello, " + "World!"); // concatenated string

// const keyword
const NISH = "Nishkarsh Raj Khare";
fmt.Println(NISH)
}
