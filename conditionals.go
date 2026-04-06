package main

import "fmt"

func main() {

	fmt.Println("Conditionals Example")

	var num1, num2 = 10, 10
	var dayOfWeek = "Saturday"

	if num1 > num2 {
		fmt.Println("num1 is greater than num2")
	} else if num1 < num2 {
		fmt.Println("num1 is less than num2")
	} else {
		fmt.Println("num1 is equal to num2")
	}

	switch dayOfWeek {
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend!")
	case "Friday":
		fmt.Println("It's almost the weekend!")
	default:
		fmt.Println("It's a regular day.")
	}

}
