package main

import (
	"fmt"
	"time"
)

type transmitter interface {
	transmitMessage()
}

type lineOne struct {
	description string
}

type lineTwo struct {
	description string
}

func (l *lineOne) transmitMessage(message string) {
	fmt.Println("Transmitting message...")
	time.Sleep(time.Duration(1))
	fmt.Println("Message transmitted via lineOne")
}

func (l *lineTwo) transmitMessage(message string) {
	fmt.Println("Transmitting message...")
	time.Sleep(time.Duration(1))
	fmt.Println("Message transmitted via lineTwo")
}

func main() {

	fmt.Println("*********Interfaces*********")

	lOne := lineOne{description: "Line One transmission line"}
	lTwo := lineTwo{description: "Line Two transmission line"}

	lOne.transmitMessage("Hello via Line One")
	lTwo.transmitMessage("Hello via Line Two")

	fmt.Println("************************************")
	fmt.Println()

}
