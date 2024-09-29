package main

import (
	"fmt"
	"math"
)

func main() {
	var radius float64
	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scan(&radius)
	luas := math.Pi * math.Pow(radius, 2)
	keliling := 2 * math.Pi * radius
	fmt.Printf("Luas lingkaran: %.2f\n", luas)
	fmt.Printf("Keliling lingkaran: %.2f\n", keliling)
}
