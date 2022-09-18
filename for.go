package main

import "fmt"

func main() {

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan Ke ", counter)
	}

	slice := []string{"Eko", "Kurniawan", "Khannedy", "Budi", "Garox"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// for loops slice using range
	for i, value := range slice {
		fmt.Println("index", i, "=", value)
	}

	person := make(map[string]string)
	person["name"] = "Eko"
	person["title"] = "Programmer"

	// for loops map using range
	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
