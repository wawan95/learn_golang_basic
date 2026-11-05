package main

import "fmt"

func getFullName() (string, string) {
	return "Sugeng", "Kurniawan"
}

func main() {
	//firstName, lastName := getFullName()
	//fmt.Println(firstName, lastName)

	firstName, _ := getFullName()
	fmt.Println(firstName)
}

/*
- Multiple return value wajib ditagkap semua valuenya
- jika ingin menghiraukan return value bisa menggunakan _
*/
