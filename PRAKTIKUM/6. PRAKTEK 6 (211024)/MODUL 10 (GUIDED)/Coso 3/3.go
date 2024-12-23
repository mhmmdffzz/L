package main

import "fmt"

func main() {
	var bilangan, d1, d2, d3, d4 int
	var teks string
	fmt.Print("bilangan : ")
	fmt.Scan(&bilangan)
	d4 = bilangan % 10
	d3 = (bilangan % 100) / 10
	d2 = (bilangan % 1000) / 100
	d1 = bilangan / 1000

	if d1 < d2 && d2 < d3 && d3 < d4 {
		teks = "Terurut Membesar"
	} else if d1 > d2 && d2 > d3 && d3 > d4 {
		teks = "Terurut Mengecil"
	} else {
		teks = "Tidak Terurut"
	}
	fmt.Println("Digit Pada Bilangan", bilangan, teks)

}
