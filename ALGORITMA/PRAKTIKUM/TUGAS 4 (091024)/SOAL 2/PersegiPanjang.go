package main

import (
	"fmt"
)

type PersegiPanjang struct {
	Panjang  float64
	Lebar    float64
	Luas     float64
	Keliling float64
}

func (p PersegiPanjang) HitungLuas() float64 {
	return p.Panjang * p.Lebar
}
func (p PersegiPanjang) HitungKeliling() float64 {
	return 2 * (p.Panjang + p.Lebar)
}
func main() {
	panjang := 10.0
	lebar := 5.0
	persegiPanjang := PersegiPanjang{
		Panjang: panjang,
		Lebar:   lebar,
	}
	persegiPanjang.Luas = persegiPanjang.HitungLuas()
	persegiPanjang.Keliling = persegiPanjang.HitungKeliling()
	fmt.Printf("Informasi Persegi Panjang:\n")
	fmt.Printf("Panjang: %.1f\n", persegiPanjang.Panjang)
	fmt.Printf("Lebar: %.1f\n", persegiPanjang.Lebar)
	fmt.Printf("Luas: %.1f\n", persegiPanjang.Luas)
	fmt.Printf("Keliling: %.1f\n", persegiPanjang.Keliling)
}
