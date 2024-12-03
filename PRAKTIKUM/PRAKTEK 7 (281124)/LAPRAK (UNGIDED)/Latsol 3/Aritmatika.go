package main

import "fmt"

func main() {
	var n, hasil int
	fmt.Scan(&n)
	switch {
	case n%10 == 0:
		hasil = n / 10
		fmt.Printf("Kategori: Bilangan Kelipatan 10\nHasil pembagian antara %d / 10 = %d\n", n, hasil)
	case n%5 == 0 && n != 5:
		hasil = n * n
		fmt.Printf("Kategori: Bilangan Kelipatan 5\nHasil kuadrat dari %d^2 = %d\n", n, hasil)
	case n%2 == 0:
		hasil = n * (n + 1)
		fmt.Printf("Kategori: Bilangan Ganjil\nHasil penjumlahan dengan bilangan %d * %d = %d\n", n, n+1, hasil)
	case n%2 != 0:
		hasil = n + (n + 1)
		fmt.Printf("Kategori: Bilangan Genap\nHasil perkalian dengan bilangan %d + %d = %d\n", n, n+1, hasil)
	default:
		fmt.Println("Tidak termasuk kategori apapun.")
	}
}
