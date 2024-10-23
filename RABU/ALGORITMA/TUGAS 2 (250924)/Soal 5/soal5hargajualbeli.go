package main

import (
	"fmt"
)

func main() {
	var hargaBeli1, hargaBeli2, hargaBeli3 int
	var hargaJual1, hargaJual2, hargaJual3 float64
	fmt.Print("Masukkan harga beli barang 1: ")
	fmt.Scan(&hargaBeli1)
	fmt.Print("Masukkan harga beli barang 2: ")
	fmt.Scan(&hargaBeli2)
	fmt.Print("Masukkan harga beli barang 3: ")
	fmt.Scan(&hargaBeli3)
	hargaJual1 = float64(hargaBeli1) * 1.05 // Keuntungan 5%
	hargaJual2 = float64(hargaBeli2) * 1.05
	hargaJual3 = float64(hargaBeli3) * 1.05
	fmt.Printf("Harga jual barang 1: %.2f\n", hargaJual1)
	fmt.Printf("Harga jual barang 2: %.2f\n", hargaJual2)
	fmt.Printf("Harga jual barang 3: %.2f\n", hargaJual3)
}
