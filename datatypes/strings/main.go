package main

import "fmt"

func main() {
	fmt.Println("*********Strings*********")

	fmt.Println("Raw String Literals. They are enclosed in back ticks [`] and prints the content as is.")
	RawStringLiteral := `This is a raw string literal with \n`
	fmt.Println(RawStringLiteral)

	fmt.Println()

	fmt.Println(`Interpreted String Literals. They are enclosed in double quotes ["] and characters can be escaped, \n will be treated as new line character.`)
	InterpretedStringLiteral := "This is a interpreted string literal with new line character at the end\n"
	fmt.Println(InterpretedStringLiteral)

	fmt.Println("************************************")
	fmt.Println()
}
