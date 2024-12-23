package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	pita := ""
	jumlah := 0
	for i := 0; i < N; i++ {
		var bunga string
		fmt.Printf("Bunga ke-%d: ", i+1)
		fmt.Scan(&bunga)
		if bunga == "selesai" {
			break
		}
		if pita == "" {
			pita += bunga
		} else {
			pita += " - " + bunga
		}
		jumlah++
	}
	if pita == "" {
		fmt.Println("Pita kosong.")
	} else {
		fmt.Printf("Pita: %s\n", pita)
		fmt.Printf("Jumlah bunga: %d\n", jumlah)
	}
}
