package main

import "fmt"

func main() {
	var gelas1, gelas2, gelas3, gelas4 string
	benar := true
	for i := 1; i <= 5; i++ {
		fmt.Scan(&gelas1, &gelas2, &gelas3, &gelas4)
		if gelas1 != "merah" || gelas2 != "kuning" ||
			gelas3 != "hijau" || gelas4 != "ungu" {
			benar = false
		}
	}
	fmt.Println(benar)
}
