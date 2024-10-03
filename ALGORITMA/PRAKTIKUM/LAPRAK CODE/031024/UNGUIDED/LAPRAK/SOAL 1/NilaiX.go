package main

import (
	"fmt"
)

func hitungX(fx float64) float64 {
	// Kurangi 5 dari f(x)
	temp := fx - 5
	// Menghitung nilai x berdasarkan persamaan
	x := (2 / temp) - 5

	return x
}

func main() {
	var fx, fx2 float64
	// Meminta input dari user
	fmt.Print("Masukkan nilai f(x)Pertama: ")
	fmt.Scan(&fx)
	fmt.Print("Masukkan nilai f(x)Kedua: ")
	fmt.Scan(&fx2)
	// Menghitung nilai x
	x := hitungX(fx)
	x1 := hitungX(fx2)
	// Menampilkan hasil
	fmt.Printf("Nilai x Pertama adalah: %.f\n", x)
	fmt.Printf("Nilai x Kedua adalah: %.f\n", x1)
}
