package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func main() {
	sayHelloWithFilter("Zaki", spamFilter)
	sayHelloWithFilter("anjing", spamFilter)
}

func spamFilter(name string) string {
	if name == "anjing	" {
		return "..."
	} else {
		return name
	}
}
