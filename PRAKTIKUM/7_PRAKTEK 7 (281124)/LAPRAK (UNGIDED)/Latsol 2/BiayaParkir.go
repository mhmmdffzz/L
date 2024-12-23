package main

import "fmt"

func main() {
	var JK string
	var waktu, tarifparkir, totalbiaya int
	fmt.Scan(&JK, &waktu)
	switch JK {
	case "Motor":
		tarifparkir = 2000
	case "Mobil":
		tarifparkir = 5000
	case "Truk":
		tarifparkir = 8000
	}
	totalbiaya = waktu * tarifparkir
	fmt.Println("Rp", totalbiaya)
}
