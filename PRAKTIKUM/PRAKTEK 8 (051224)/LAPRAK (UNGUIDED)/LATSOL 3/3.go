package main

import "fmt"

func main() {
	var angka, pembagi, hasil int
	fmt.Scan(&angka, &pembagi)
	for angka >= pembagi {
		angka = angka - pembagi
		hasil++
	}
	fmt.Print(hasil)
}
