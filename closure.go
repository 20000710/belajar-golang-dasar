package main

import "fmt"

func main() {
	name := "Eko"
	counter := 0

	increment := func() {
		name = "khairi"
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
