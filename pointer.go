package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1 // same with address2 := &address1
	var address3 *Address = &address1

	address2.City = "Bandung"

	// any variable who reference to address 1 will get a result from *address2
	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	// make new pointer reference to struct Address
	var address4 *Address = new(Address)
	// var address4 *Address = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	address4.City = "Jakarta"
	address4.Province = "DKI Jakarta"
	address4.Country = "Indonesia"
	fmt.Println(address4)
}
