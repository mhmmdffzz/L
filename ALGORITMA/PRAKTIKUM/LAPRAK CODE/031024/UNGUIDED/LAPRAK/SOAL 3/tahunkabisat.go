package main

import "fmt"

func main() {
	var tahun int
	// input Umur Kabisat
	fmt.Print("Tahun :")
	fmt.Scan(&tahun)
	if (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0) {
		fmt.Println("kabisat : True")
	} else {
		fmt.Println("kabisat : False")
	}

}
