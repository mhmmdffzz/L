package main

import "fmt"

func main() {
	var bilangan int
	fmt.Print("Masukkan bilangan tiga digit: ")
	fmt.Scan(&bilangan)

	d1, d2, d3 := bilangan/100, (bilangan/10)%10, bilangan%10
	fmt.Println(d1 > d2 && d2 > d3)
}
