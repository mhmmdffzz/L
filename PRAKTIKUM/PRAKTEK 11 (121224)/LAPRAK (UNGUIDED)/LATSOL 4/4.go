package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	sum := 0.0
	var i int
	for i = 0; i < n; i++ {
		term := 1.0 / float64(2*i+1)
		if i%2 != 0 {
			term = -term
		}
		sum += term
		pi := 4 * sum

		nextTerm := 1.0 / float64(2*(i+1)+1)
		if (i+1)%2 != 0 {
			nextTerm = -nextTerm
		}

		if math.Abs(nextTerm) < 0.00001 {
			break
		}

		if pi >= 3.1415876535 {
			fmt.Printf("Hasil PI: %.10f\n", pi)
		}
	}
	fmt.Printf("Pada i ke: %d\n", i)
}
