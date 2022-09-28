package main

import "fmt"

func getFullname2() (firstName string, middleName string, lastName string) {
	firstName = "Zaki"
	middleName = "Khairi"
	lastName = "Ziwar"
	return
}

func main() {
	firstName, middleName, lastName := getFullname2()
	fmt.Println(firstName, middleName, lastName)
}
