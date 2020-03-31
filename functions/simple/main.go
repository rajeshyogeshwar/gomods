package main

import "fmt"

func multiply(first int, second int) int {
	return first * second
}

func divmod(first int, second int) (int, int) {
	return first / second, first % second
}

func main() {
	product := multiply(10, 15)
	fmt.Printf("Product of 10 and 15 is %d which was calculated using multiply method\n", product)

	fmt.Println()

	quotient, remainder := divmod(15, 2)
	fmt.Println("Method returning multiple values")
	fmt.Printf("Quotient after dividing 15 by 2 is %d and remainder is %d which was calculated using divmod method\n", quotient, remainder)

}
