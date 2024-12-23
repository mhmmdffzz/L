package main

import "fmt"

func main() {
	var desimal, bilbulat float64
	fmt.Scan(&desimal)
	bilbulat = float64(int64(desimal)) + 1.0
	for {
		desimal = desimal + 0.1
		if desimal >= bilbulat-0.0000001 {
			break
		}
		fmt.Printf("%0.1f\n", desimal)
	}
	fmt.Print(int64(bilbulat))
}
