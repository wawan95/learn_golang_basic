package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Sugeng"
	middleName = "Kurniawan"
	lastName = "Developer"

	return firstName, middleName, lastName
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}
