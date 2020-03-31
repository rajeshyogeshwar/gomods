package main

import "fmt"

func main() {
	fmt.Println("*********Arrays*********")

	StringArray := [5]string{"Rajesh", "Akshay", "Mukesh", "Saurabh", "Pradeep"}
	fmt.Printf("Type %T of StringArray variable\n", StringArray)
	fmt.Printf("Value of StringArray %v\n", StringArray)

	fmt.Println("Looping over the items and printing it")
	for i := 0; i < len(StringArray); i++ {
		fmt.Println(StringArray[i])
	}

	var TwoDimensionalIntArray [5][2]int
	for i := 0; i < 5; i++ {
		for j := 0; j < 2; j++ {
			TwoDimensionalIntArray[i][j] = (i + 1) * (j + 1)
		}
	}

	fmt.Println()
	fmt.Printf("Two Dimensional Array %v\n", TwoDimensionalIntArray)

	fmt.Println("************************************")
	fmt.Println()
}
