package main

import (
	"fmt"
	"math"
)

func changeTheNumber(number *int) {

	*number = int(math.Pow(float64(*number), 3))

}

func modifySlice(slice *[]int) {
	*slice = append(*slice, 4)
	*slice = append(*slice, 5)
}

func main() {

	fmt.Println("*********Pointers*********")

	fmt.Println("A simple variable number with initial value of 15 is passed to a function accept integer pointer")
	number := 15
	changeTheNumber(&number)
	fmt.Println(number)

	fmt.Println()

	fmt.Println("A slice is passed to a function accepting a slice pointer where 2 more values are appended to the slice")
	slice := []int{1, 2, 3}
	modifySlice(&slice)
	fmt.Println(slice)

	fmt.Println("************************************")
	fmt.Println()

}
