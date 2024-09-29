package main

import (
	"fmt"
	"math"
)

func main() {
	var jariJari float64

	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scanln(&jariJari)

	// Menghitung luas lingkaran
	luas := math.Pi * math.Pow(jariJari, 2)

	// Menampilkan hasil perhitungan luas lingkaran
	fmt.Printf("Luas lingkaran dengan jari-jari %.2f adalah %.2f\n", jariJari, luas)
}
