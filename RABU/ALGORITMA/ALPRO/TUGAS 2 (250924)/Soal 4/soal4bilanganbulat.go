package main

import (
	"fmt"
)

func main() {
	var x int
	var d1, d2, d3 int
	fmt.Print("Masukkan bilangan bulat positif x (<= 999): ")
	fmt.Scan(&x)
	if x < 0 || x > 999 {
		fmt.Println("Input tidak valid. Pastikan x adalah bilangan bulat positif <= 999.")
		return
	}

	d3 = x % 10         // Digit satuan
	d2 = (x / 10) % 10  // Digit puluhan
	d1 = (x / 100) % 10 // Digit ratusan

	fmt.Printf("Digit pertama (ratusan): %d\n", d1)
	fmt.Printf("Digit kedua (puluhan): %d\n", d2)
	fmt.Printf("Digit ketiga (satuan): %d\n", d3)
}
