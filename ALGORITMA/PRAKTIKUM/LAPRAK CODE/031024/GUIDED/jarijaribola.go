package main

import (
	"fmt"
	"math"
)

func main() {
	var radius int
	// input jari-jari dari user
	fmt.Print("Jejari = ")
	fmt.Scan(&radius)
	// nilai pi
	pi := 3.1415926535
	// Menghitung volume bola
	volume := (4.0 / 3.0) * pi * math.Pow(float64(radius), 3)
	// Menghitung luas permukaan bola
	luas := 4 * pi * math.Pow(float64(radius), 2)
	// hasil
	fmt.Printf("Bola dengan jejari %d memiliki volume %.4f dan luas kulit %.4f\n", radius, volume, luas)
}
