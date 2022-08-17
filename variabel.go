package main

import "fmt"

func main() {

	// declare variabel with data types
	var name string

	name = "Zaki Khairi"
	fmt.Println(name)

	name = "Zaki Zwar"
	fmt.Println(name)

	// declare variabel with no data types
	var friendName = "Karom"
	fmt.Println(friendName)

	var age = 30
	fmt.Println(age)

	// declare variabel with :=
	country := "Indonesia"
	fmt.Println(country)

	country = "USA"

	//declare multiple variabel
	var (
		firstname = "Zaki"
		lastname  = "KhairiZiwar"
	)

	fmt.Println(firstname, lastname)
}
