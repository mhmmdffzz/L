package main

import "fmt"

func main() {
	var tahun1, tahun2, tahun3 int
	fmt.Print("Tahun1 :")
	fmt.Scan(&tahun1)
	fmt.Print("Tahun2 :")
	fmt.Scan(&tahun2)
	fmt.Print("Tahun3 :")
	fmt.Scan(&tahun3)
	if (tahun1%400 == 0) || (tahun1%4 == 0 && tahun1%100 != 0) {
		fmt.Println("kabisat1 : True")
	} else {
		fmt.Println("kabisat1 : False")
	}
	if (tahun2%400 == 0) || (tahun2%4 == 0 && tahun2%100 != 0) {
		fmt.Println("kabisat2 : True")
	} else {
		fmt.Println("kabisat2 : False")
	}
	if (tahun3%400 == 0) || (tahun3%4 == 0 && tahun3%2100 != 0) {
		fmt.Println("kabisat3 : True")
	} else {
		fmt.Println("kabisat3 : False")
	}
}
