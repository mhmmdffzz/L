package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n < 0 {
		fmt.Print("False, harus bilangan positif")
		return
	}
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	fmt.Println(hasil)
}
