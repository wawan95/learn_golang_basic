package main

import "fmt"

func getHello(name string) string {
	hello := "Hello " + name
	return hello
}

func main() {
	result := getHello("Sugeng")
	fmt.Println(result)

	fmt.Println(getHello("Nico"))
	fmt.Println(getHello("Silvi"))
}
