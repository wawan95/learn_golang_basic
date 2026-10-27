package main

import "fmt"

func main() {
	name := "Sugeng"

	if name == "Sugeng" {
		fmt.Println("Hello Sugeng")
	} else if name == "Ridwan" {
		fmt.Println("Hello Ridwan")
	} else if name == "Joko" {
		fmt.Println("Hello Joko")
	} else {
		fmt.Println("Hi, Boleh Kenalan?")
	}

	if length := len(name); length > 6 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
