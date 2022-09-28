package main

import "fmt"

func getFullName() (string, string, string) {
	return "Zaki", "Khairi", "Ziwar"
}

func main() {
	firstName, middleName, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}
