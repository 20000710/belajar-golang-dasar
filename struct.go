package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
	// Married       bool
}

// struct method
func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main() {
	var cust Customer
	cust.Name = "Eko"
	cust.Address = "Indonesia"
	cust.Age = 30

	// call struct method
	cust.sayHi("Abdul")

	// fmt.Println(cust.Name)
	// fmt.Println(cust.Address)
	// fmt.Println(cust.Age)

	// struct literals
	// joko := Customer{
	// 	Name:    "Joko",
	// 	Address: "Cirebon",
	// 	Age:     30,
	// }
	// fmt.Println(joko)

	// struct literals without field declaration
	// budi := Customer{"Budi", "Jakarta", 35}
	// fmt.Println(budi)
}
