package main

import "fmt"

func main() {

	fmt.Println("*********Ranges*********")

	NumericArray := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("Using range over an integer array")
	for _, number := range NumericArray {
		fmt.Println(number)
	}

	fmt.Println()

	fmt.Println("Using range over a map to print it's keys and corresponding values")
	UserDataMap := map[string]string{"name": "Rajesh Yogeshwar", "profession": "Software Engineer", "location": "Mumbai, India"}

	for key, value := range UserDataMap {
		fmt.Println(key, value)
	}

	fmt.Println()

	fmt.Println("Using range over a string where it prints index and character code for character at that index. Eg. if character at 0 is a, then output is 0 97")
	for index, character := range "Rajesh Yogeshwar" {
		fmt.Println(index, character)
	}

}
