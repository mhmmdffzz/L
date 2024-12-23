package main

import (
	"fmt"
)

func hitungFx(x float64) float64 {
	return (2 / (x + 5)) + 5
}
func main() {
	var x1, x2 float64
	fmt.Print("Masukkan nilai x Pertama: ")
	fmt.Scan(&x1)
	fmt.Print("Masukkan nilai x Kedua: ")
	fmt.Scan(&x2)
	fx1 := hitungFx(x1)
	fx2 := hitungFx(x2)
	fmt.Printf("Nilai f(x) untuk x Pertama (%.1f) adalah: %.1f\n", x1, fx1)
	fmt.Printf("Nilai f(x) untuk x Kedua (%.1f) adalah: %.3f\n", x2, fx2)
}
