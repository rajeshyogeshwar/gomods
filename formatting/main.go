package main

import (
	"fmt"
	"os"
)

/* Demonstration of different types of print/format methods that are available in Golang to print values*/

func divideNumber(numberOne int, numberTwo int) int {
	return numberOne / numberTwo
}

func main() {

	numberOne := 12
	numberTwo := 2

	fmt.Println("Demo for methods from fmt package")

	fmt.Println("Demo for Printf method which prints the string as per the format specifiers used. This is very similar to how it was in C. \nPrintf doesn't append newline character at the end of string")
	fmt.Printf("Quotient is %d ", divideNumber(numberOne, numberTwo))

	fmt.Println()
	fmt.Println("********")
	fmt.Println()

	fmt.Println("This line is printed using Println")

	fmt.Println()
	fmt.Println("********")
	fmt.Println()

	fmt.Fprint(os.Stderr, "This is line printed using Fprint. The first parameter to this method is any writer. I have Stderr of os package")

	fmt.Println()
	fmt.Println("********")
	fmt.Println()

	fmt.Fprintf(os.Stderr, "This is line printed using Fprintf. The first parameter to this method is any writer. I have Stderr of os package. But is also allows for formatting. Like including %d in here.", 4)

	fmt.Println()
	fmt.Println("********")
	fmt.Println()

	fmt.Printf("Simple usage of Printf with formatting available. Just like how I am inserting %s using format specifier", "Jackass")

	fmt.Println()

	fmt.Println("We have already made use of Println which adds a newline character at the end of the content being printed")

}
