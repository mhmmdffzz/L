package main

import (
	"fmt"
	"math"
)

func main() {
	var radius int
	fmt.Print("Jejari = ")
	fmt.Scan(&radius)
	pi := 3.1415926535
	volume := (4.0 / 3.0) * pi * math.Pow(float64(radius), 3)
	luas := 4 * pi * math.Pow(float64(radius), 2)
	fmt.Printf("Bola dengan jejari %d memiliki volume %.4f dan luas kulit %.4f\n", radius, volume, luas)
}
