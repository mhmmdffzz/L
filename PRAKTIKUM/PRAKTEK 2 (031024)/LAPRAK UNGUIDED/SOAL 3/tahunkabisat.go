package main

import "fmt"

func main() {
	var tahun int
	fmt.Print("Tahun :")
	fmt.Scan(&tahun)
	if !(tahun%400 == 0) && !(tahun%4 == 0 && tahun%100 != 0) {
		fmt.Println("kabisat : False")
	} else {
		fmt.Println("kabisat : True")
	}
}
