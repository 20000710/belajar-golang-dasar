package main

import "fmt"

func main() {
	name := "Garox"

	if name == "Zaki" {
		fmt.Println("Halo Zaki")
	} else if name == "Garox" {
		fmt.Println("Halo Garox")
	} else {
		fmt.Println("Hi, Boleh kenalan?")
	}

	// short if statement
	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
