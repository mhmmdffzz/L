package main

import (
	"fmt"
)

func main() {
	var x1, x2 float64
	fmt.Print("Masukkan angka pertama: ")
	fmt.Scan(&x1)
	fmt.Print("Masukkan angka kedua: ")
	fmt.Scan(&x2)
	fx1 := (2 / (x1 + 5)) + 5
	fx2 := (2 / (x2 + 5)) + 5
	fmt.Printf("Hasil f(x) untuk angka pertama: %.1f\n", fx1)
	fmt.Printf("Hasil f(x) untuk angka kedua: %.3f\n", fx2)
}
