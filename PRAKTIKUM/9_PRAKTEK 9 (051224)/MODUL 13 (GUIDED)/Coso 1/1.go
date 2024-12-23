package main

import "fmt"

func main() {
	var usr, pwd string
	var kondisi bool
	for kondisi = false; !kondisi; {
		fmt.Scan(&usr, &pwd)
		kondisi = usr == "admin" && pwd == "admin12345"
	}
	fmt.Println("Selamat, Anda berhasil login ")
}
