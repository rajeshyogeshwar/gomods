package main

import "fmt"

func main() {
	fmt.Println("*********Maps*********")

	UserDataMap := map[string]string{"name": "Rajesh Yogeshwar", "profession": "Software Engineer", "location": "Mumbai, India"}
	fmt.Printf("Type %T\n", UserDataMap)
	fmt.Printf("Value %v\n", UserDataMap)

	fmt.Println()
	fmt.Printf("Name is %s\n", UserDataMap["name"])
	UserDataMap["age"] = "31"

	fmt.Printf("Map after adding age key %v\n", UserDataMap)
	fmt.Println()

	fmt.Println("Checking presence of dateOfBirth key using _, present where _ is a blank identifier and the value is ignored")
	_, present := UserDataMap["dateOfBirth"]
	fmt.Printf("dateOfBirth is present %v\n", present)

	delete(UserDataMap, "age")
	fmt.Printf("Map after removing age using delete method key %v\n", UserDataMap)
	fmt.Println()

	fmt.Println("************************************")
	fmt.Println()
}
