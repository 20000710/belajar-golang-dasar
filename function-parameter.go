package main

import "fmt"

func sayHelloTo(firstname string, lastname string) {
	fmt.Println("Hello", firstname, lastname)
}

func main() {
	firstName := "Zaki"
	sayHelloTo(firstName, "Khairi")
	sayHelloTo("Budi", "Nugraha")
}
