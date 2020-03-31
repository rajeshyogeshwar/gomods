package main

import "fmt"

func sequence() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func makeFibGen() func() int {
	f1 := 0
	f2 := 1
	return func() int {
		f2, f1 = (f1 + f2), f2
		return f1
	}
}

func main() {

	// What happens here is closedFunction actually gets assigned the closed function which has a copy of i with itself and keeps on iterating it. Hence every time we call it, the value increments. But when we again assign sequence to another variable the value resets. Same holds true for makeFibGen too. The values in comments are indicative of this.

	fmt.Println("Sequence")

	closedFunction := sequence()
	fmt.Println(closedFunction())
	fmt.Println(closedFunction())
	fmt.Println(closedFunction())
	fmt.Println(closedFunction())

	fmt.Println()
	fmt.Println("Fibonacci")
	fmt.Println()

	generator := makeFibGen()
	fmt.Println(generator()) // f2 is 1, f1 is 1
	fmt.Println(generator()) // f2 is 2, f1 is 1
	fmt.Println(generator()) // f2 is 3, f1 is 2
	fmt.Println(generator()) // f2 is 5, f1 is 3
	fmt.Println(generator()) // f2 is 8, f1 is 5
	fmt.Println(generator()) // f2 is 13, f1 is 8

}
