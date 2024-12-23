package main

import (
	"fmt"
)

func main() {
	var berat1, berat2 float64
	fmt.Scan(&berat1, &berat2)
	if berat1 < 0 || berat2 < 0 {
		fmt.Println("Proses input selesai karena salah satu kantong beratnya negatif.")
		return
	}
	totalBerat := berat1 + berat2
	if totalBerat > 150 {
		fmt.Println("Proses input selesai karena total berat isi kedua kantong melebihi 150 kg.")
		return
	}
	selisih := berat1 - berat2
	if selisih < 0 {
		selisih = -selisih
	}
	if selisih >= 9 {
		fmt.Println("sepeda motor pak andi akan oleng : true")
	} else {
		fmt.Println("sepeda motor pak andi akan oleng : false")
	}
}
