package main

import "fmt"

func main() {
	var angka int
	fmt.Scan(&angka)
	for angka > 0 {
		fmt.Println(angka % 10)
		angka = angka / 10
	}
}
