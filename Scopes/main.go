//in same folder package name will be same
package main 
import "fmt"
/* multiline comment */
var x int = 10 // package + local scoped variable, accessible throughout the package

func main() {
	result := add(x, 5) // Add is package scoped function.
	// Result is blocked scoped variable, only accessible within main function.
	fmt.Println("The result is:", result)
}