package main

import "fmt"

func main() {
	var usia int
	var kk bool
	fmt.Scan(&usia, &kk)
	if usia >= 17 && kk {
		fmt.Println("Bisa Membuat KTP")
	} else {
		fmt.Println("Belum Bisa Membuat KTP")
	}

}
