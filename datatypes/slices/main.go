package main

import "fmt"

func modifySlice(slice []int) {
	slice = append(slice, 99)
	slice = append(slice, 999)
	slice = append(slice, 9999)
}

func main() {
	fmt.Println("*********Slices*********")
	stringSlice := []string{"Rajesh", "Akshay", "Mukesh", "Saurabh", "Pradeep"}

	fmt.Printf("Type %T Original Slice %v", stringSlice, stringSlice)
	stringSlice = append(stringSlice, "Pratik")
	fmt.Printf("After appending Pratik to slice it is now %v", stringSlice)

	stringSlice = append(stringSlice, "Rupesh")
	fmt.Printf("After appending Rupesh to slice it is now %v", stringSlice)

	fmt.Println("Printing elements at 1st and second position using :")
	fmt.Println(stringSlice[1:3])

	fmt.Println("************************************")
	fmt.Println()

}
