package main

import "fmt"

func main() {
	var diskon, totalBelanja, totalAkhir int
	fmt.Print("masukkan total belanja: ")
	fmt.Scan(&totalBelanja)
	fmt.Print("masukkan diskon (%): ")
	fmt.Scan(&diskon)
	totalAkhir = totalBelanja - (totalBelanja * diskon / 100)
	fmt.Printf("total belanja akhir setelah diskon: %d\n", totalAkhir)

}
