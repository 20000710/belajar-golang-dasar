package main

import "fmt"

func main() {

	name := "Abdul"

	switch name {
	case "Zaki":
		fmt.Println("Halo, Zaki")
	case "Garox":
		fmt.Println("Halo, Garox")
	default:
		fmt.Println("Boleh berkenalan?")
	}

	// short switch case statement
	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("Nama terlalu panjang")
	// case false:
	// 	fmt.Println("Nama sudah benar")
	// }

	// simplify switch case
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length < 5:
		fmt.Println("Nama terlalu pendek")
	default:
		fmt.Println("Nama sudah benar")
	}
}
