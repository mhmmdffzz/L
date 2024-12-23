package main

import "fmt"

func main() {
	var K int
	fmt.Scan(&K)
	if K < 0 {
		fmt.Println("Nilai K harus bilangan positif!")
	} else {
		fmt.Printf("Nilai K = %d\n", K)
		hasil := 1.0
		for k := 0; k <= K; k++ {
			pembilang := (4*k + 2) * (4*k + 2)
			penyebut := (4*k + 1) * (4*k + 3)
			hasil *= float64(pembilang) / float64(penyebut)
		}
		fmt.Printf("Nilai akar 2 = %.10f\n", hasil)
	}
}
