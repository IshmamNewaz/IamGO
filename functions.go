package main

import "fmt"

func divide(a, b int /* Both are integers */) (float64, int /* Both returns are integers */) {
	return float64(a) / float64(b), a % b //Need to typecast to float64 for division to get a float result, and modulus operator % is used to get the remainder of the division
}

//will return the quotient and remainder of the multiplication of a and b.
func multiply(a, b int) (quotient int, remainder int) {
	quotient = a * b
	remainder = 0
	return
}

func subs(a, b int) int {
	return a - b
}

func main() {
	fmt.Println("Functions Example")

	var a, b = 15, 4

	c, d := divide(a, b)
	fmt.Printf("Quotient: %.2f, Remainder: %d\n", c, d) // %f is used for float	placeholder and %d is used for integer placeholder

	e, f := multiply(a, b)
	fmt.Printf("Quotient: %d, Remainder: %d\n", e, f)

	//The "annonymousFunc" is a variable that holds an anonymous function. It holds the function but does not automatically execute it. To execute the function, we need to call it with the appropriate arguments.
	annonymousFunc := func(x int) int { // This is an anonymous function that is allowed to be written inside another function and can be assigned to a variable.
		return x * x
	}
	fmt.Printf("Square of %d is %d\n", a, annonymousFunc(a))

	//This is an example of an immediately invoked function expression (IIFE) in Go. It defines an anonymous function that takes an integer argument and returns its square, and then immediately invokes it with the argument 5. The result is stored in the variable "result" and printed to the console.
	result := func(x int) int {
		return x * x
	}(5)

	fmt.Printf("Square of 5 is %d\n", result)

	//Functions as a parameter

	mynumber1 := 10
	mynumber2 := 20
	mynumber3 := 30

	//The "applyOperation" function is a higher-order function that takes two integers and a function as parameters.
	// operation is a variable which stores the function itself

	applyOperation := func(x, y int, operation func(int, int) int) int {
		return operation(x, y)
	}
	sum := applyOperation(mynumber1, mynumber2, func(x, y int) int {
		return x + y + mynumber3
	})
	fmt.Printf("Sum of %d and %d is %d\n", mynumber1, mynumber2, sum)

	applyOperation2 := func(x, y int, substraction func(int, int) int) int {
		return substraction(x, y)

	}
	sum2 := applyOperation2(mynumber1, mynumber2, subs) // Here we are passing the "subs" function as an argument to the "applyOperation2" function. The "subs" function is defined earlier in the code and takes two integers as input and returns their difference. By passing the "subs" function as an argument, we can use it within the "applyOperation2" function to perform the subtraction operation on the provided integers a and b.
	fmt.Printf("Difference of %d and %d is %d\n", mynumber1, mynumber2, sum2)

}
