package main

import "fmt"

func main() {
	var b1, b2, b3, max, min int
	fmt.Scan(&b1, &b2, &b3)

	if b1 > b2 {
		max = b1
		min = b2
	} else {
		max = b2
		min = b1
	}
	if max < b3 {
		max = b3
	}
	if min > b3 {
		min = b3
	}
	fmt.Println("Terbesar", max)
	fmt.Println("Terkecil", min)
}
