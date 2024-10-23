package main

import (
	"fmt"
)

func main() {
	var x, y int
	var hasil float64
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)
	hasil = 1/(float64(3*x*x)+10) + float64(10*y) + 7
	fmt.Printf("Nilai dari f(%d, %d) adalah: %.2f\n", x, y, hasil)
}
