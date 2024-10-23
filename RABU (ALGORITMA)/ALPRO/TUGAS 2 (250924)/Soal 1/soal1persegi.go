package main

import "fmt"

func main() {
	var p, l int

	fmt.Print("Masukkan panjang persegi panjang: ")
	fmt.Scan(&p)
	fmt.Print("Masukkan lebar persegi panjang: ")
	fmt.Scan(&l)

	luas := p * l
	keliling := 2 * (p + l)

	fmt.Println("Luas persegi panjang:", luas)
	fmt.Println("Keliling persegi panjang:", keliling)
}
