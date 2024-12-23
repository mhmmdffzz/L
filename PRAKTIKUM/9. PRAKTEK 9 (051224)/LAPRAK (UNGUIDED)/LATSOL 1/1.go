package main

import "fmt"

func main() {
	var angka, jumlah int64
	jumlah = 0
	fmt.Scan(&angka)
	for angka > 0 {
		angka = angka / 10
		jumlah++
	}
	fmt.Println("banyaknya digit:", jumlah)
}
