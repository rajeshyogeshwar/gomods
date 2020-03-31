package main

import (
	"fmt"
	"math"
)

func factorial(number int) int {
	if number == 1 {
		return 1
	}
	return number * factorial(number-1)
}

func palindromeChecker(number int, newNumber int, digits []int) bool {

	quotient := newNumber / 10
	remainder := newNumber % 10
	digits = append(digits, remainder)

	if newNumber < 10 {

		lengthOfDigits := len(digits)
		power := lengthOfDigits - 1
		result := 0

		for _, digit := range digits {
			exponential := math.Pow10(power)
			result = result + (int(exponential) * digit)
			power = power - 1
		}

		if result == number {
			return true
		}
		return false

	}

	return palindromeChecker(number, quotient, digits)

}

func isNumberPalindrome(number int) bool {

	if number > 0 && number < 10 {
		return false
	}

	digits := []int{}
	return palindromeChecker(number, number, digits)

}

func main() {

	result := factorial(5)
	fmt.Printf("Factorial is %d\n", result)

	fmt.Println()
	fmt.Println("Palindrome test using recursion")
	numbers := []int{131, 9765115679, 456217845, 8596336958}

	for _, number := range numbers {

		palindromeResult := isNumberPalindrome(number)

		if palindromeResult {
			fmt.Printf("%d is a palindrome number\n", number)
		} else {
			fmt.Printf("%d is not a palindrome number\n", number)
		}

	}

}
