package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n <= 0 {
		return
	}
	count := 0
	for i := 1; i <= n; i++ {
		if i%2 != 0 {
			count++
		}
	}
	fmt.Printf("Terdapat %d bilangan ganjil\n", count)
}
