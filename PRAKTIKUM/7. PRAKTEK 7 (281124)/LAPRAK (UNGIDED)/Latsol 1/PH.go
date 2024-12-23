package main

import "fmt"

func main() {
	var PH float64
	fmt.Scan(&PH)
	switch {
	case PH >= 6.5 && PH <= 8.6:
		fmt.Println("Air Layak Minum")
	case PH < 0 || PH > 14:
		fmt.Println("Nilai PH Tidak Valid")
	default:
		fmt.Println("Air Tidak Layak Minum")
	}
}
