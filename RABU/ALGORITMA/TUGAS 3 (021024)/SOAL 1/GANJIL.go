package main

import "fmt"

func main() {
	var bilangan int
	var ganjil bool

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&bilangan)
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&bilangan)

	ganjil = (bilangan%2 != 1)
	fmt.Println(ganjil)
	ganjil = (bilangan%2 != 0)
	fmt.Println(ganjil)
}
