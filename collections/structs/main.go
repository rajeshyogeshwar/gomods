package main

import "fmt"

type userData struct {
	firstName string
	lastName  string
	email     string
	age       float64
	address   map[string]string
}

func createUserData(firstName string, lastName string, email string, age float64, address map[string]string) userData {

	newUserData := userData{firstName: firstName, lastName: lastName, email: email, age: age, address: address}

	return newUserData

}

func printUserData(users []userData) {
	for _, user := range users {
		fmt.Printf("%v\n", user)
	}
}

func main() {

	fmt.Println("*********Structs*********")
	address := map[string]string{"firstLine": "Peaceville Apts", "city": "Serene City"}
	users := []userData{}

	users = append(users, createUserData("Rajesh", "Yogeshwar", "rajesh.yogeshwar@mailinator.com", 31.0, address))

	users = append(users, createUserData("Jackal", "Jackson", "jackal.jackson@mailinator.com", 31.0, address))

	users = append(users, createUserData("Peter", "Shawn", "peter.shawn@mailinator.com", 31.0, address))

	printUserData(users)

	fmt.Println("************************************")
	fmt.Println()

}
