package main

import "fmt"

func main() {
	var mk string = "Algoritma dan Pemrograman"
	var kode string = "CAK123"
	var sks int = 3

	fmt.Print("Tuliskan kode MK dan SKS: ")
	fmt.Scan(&kode, &sks)
	fmt.Println("Kredit MK", kode, "-", mk, "1 adalah", sks, "SKS")
}
