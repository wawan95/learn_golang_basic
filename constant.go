package main

import "fmt"

func main() {
	const (
		firstName string = "Sugeng"
		lastName         = "Kurniawan"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	// error
	// firstName = "Budi"
	// lastName = "Joko"
}
