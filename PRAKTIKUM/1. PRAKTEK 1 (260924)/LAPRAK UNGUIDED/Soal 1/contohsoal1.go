package main

import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp            string
	)

	fmt.Print("Masukan input string pertama: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string kedua: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string ketiga: ")
	fmt.Scanln(&tiga)

	fmt.Println("Output awal =", satu, dua, tiga)

	// Proses pertukaran
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp

	fmt.Println("Output akhir =", satu, dua, tiga)
}
