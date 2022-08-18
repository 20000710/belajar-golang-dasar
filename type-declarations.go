package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKTPUser NoKTP = "31751012321424"
	var marriedStatus = true
	fmt.Println(noKTPUser)
	fmt.Println(marriedStatus)
}
