package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	var x int
	x = 10
	fmt.Println("Value of x:", x)

	var y float64 = 3.14
	fmt.Println("Value of y:", y)

	//Multiple variable declaration
	var a, b, c = 1, "Go", true
	fmt.Println("Values of a, b, c:", a, b, c)

	//blocked declaration
	var (
		d int    = 20
		e string = "Programming"
		f bool   = false
	)
	fmt.Println("Values of d, e, f:", d, e, f)

	//Short variable declaration
	p, q := 42, "Hello"
	fmt.Println("Values of p and q:", p, q)
}
