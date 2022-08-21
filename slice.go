package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}
	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// ubah array index ke 5
	// months[5] = "Diubah"
	// fmt.Println(slice1)

	// mengubah slice index 0 yaitu Mei menjadi Mei Lagi
	// slice1[0] = "Mei Lagi"
	// fmt.Println(months)

	var slice2 = months[2:4]
	fmt.Println(slice2)

	// append khair ke slice2 dan akan membuat array baru karena length > capacity
	// jika capacity > length maka yang terkena manipulasi adalah array lama
	var slice3 = append(slice2, "Khair")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	// membuat slice dari awal make([]type, length, capacity)
	newSlice := make([]string, 2, 5)

	newSlice[0] = "Zaki"
	newSlice[1] = "Khairi"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// membuat copy slice copy(toSlice, fromSlice)
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	// perbedaan slice dan array
	iniArray := [...]int{1, 2, 3, 4, 5} // [...] atau [length array]
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
