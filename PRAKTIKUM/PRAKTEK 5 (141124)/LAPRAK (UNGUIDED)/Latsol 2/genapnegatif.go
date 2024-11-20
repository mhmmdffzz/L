package main

import "fmt"

func main() {
	var k int
	fmt.Scan(&k)
	if k%2 == 0 && k < 0 {
		fmt.Println("genap negatif")
	} else {
		fmt.Println("bukan")
	}
}
