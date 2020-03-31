package main

import "fmt"

func variadicSumFunction(nums ...int) int {
	sum := 0
	for _, number := range nums {
		sum += number
	}
	return sum
}

func main() {
	sum := variadicSumFunction(10, 15, 25)
	fmt.Println(sum)
}
