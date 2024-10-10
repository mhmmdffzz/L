package main

import (
	"fmt"
)

type persegipanjang struct {
	panjang  float64
	lebar    float64
	luas     float64
	keliling float64
}

func main() {
	var t persegipanjang
	fmt.Println("Informasi panjang persegi panjang:")

	fmt.Print("panjang: ")
	fmt.Scanln(&t.panjang)

	fmt.Print("lebar: ")
	fmt.Scanln(&t.lebar)

	t.luas = t.panjang * t.lebar
	t.keliling = 2 * (t.panjang + t.lebar)

	fmt.Println("Luas persegi panjang:", t.luas)
	fmt.Println("Keliling persegi panjang:", t.keliling)
}
