package main

import (
	"fmt"

	funclibrary "example.com/funcLibrary" // Importing the funcLibrary package
)

func main() {
	result := funclibrary.Add(10, 7) // Using the Add function from the funcLibrary package
	fmt.Println("The result is:", result)
}
