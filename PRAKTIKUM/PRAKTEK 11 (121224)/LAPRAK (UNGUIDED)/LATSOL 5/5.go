package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var n int
	fmt.Print("Banyak Topping: ")
	fmt.Scan(&n)
	topping := 0
	centerX, centerY := 0.5, 0.5
	radius := 0.5

	for i := 0; i < n; i++ {
		x := rand.Float64()
		y := rand.Float64()
		dx := x - centerX
		dy := y - centerY
		if dx*dx+dy*dy <= radius*radius {
			topping++
		}
	}
	fmt.Printf("Topping pada Pizza: %d\n", topping)
	fmt.Printf("PI : %.10f\n", 4.0*float64(topping)/float64(n))
}
