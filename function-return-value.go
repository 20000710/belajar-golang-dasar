package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello Bro"
	} else {
		return "Hello " + name
	}

	result := "Hello"
	return result
}

func main() {
	result := getHello("Zaki")
	fmt.Println(result)
	fmt.Println(getHello(""))
}
