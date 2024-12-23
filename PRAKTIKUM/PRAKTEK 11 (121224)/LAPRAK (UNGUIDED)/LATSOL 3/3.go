package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var drops int
	fmt.Print("Masukkan jumlah tetesan air: ")
	fmt.Scan(&drops)
	countA, countB, countC, countD := 0, 0, 0, 0
	for i := 0; i < drops; i++ {
		x := rand.Float64()
		y := rand.Float64()
		if x < 0.5 {
			if y < 0.5 {
				countA++
			} else {
				countD++
			}
		} else {
			if y < 0.5 {
				countB++
			} else {
				countC++
			}
		}
	}
	fmt.Printf("Curah hujan daerah A: %.4f mm\n", float64(countA)*0.0001)
	fmt.Printf("Curah hujan daerah B: %.4f mm\n", float64(countB)*0.0001)
	fmt.Printf("Curah hujan daerah C: %.4f mm\n", float64(countC)*0.0001)
	fmt.Printf("Curah hujan daerah D: %.4f mm\n", float64(countD)*0.0001)
}
