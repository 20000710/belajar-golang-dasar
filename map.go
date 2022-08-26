package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "Zaki",
		"address": "Jakarta",
	}

	// add key title with value Programmer
	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Zaki Khairi"
	book["ups"] = "Salah"

	// map book with key ups before deleted
	fmt.Println("book: ", book)
	fmt.Println("length: ", len(book))

	delete(book, "ups")

	// map book with key ups after deleted
	fmt.Println("book: ", book)
	fmt.Println("length: ", len(book))
}
