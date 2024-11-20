package main

import "fmt"

func main() {
	var motor int
	fmt.Scan(&motor)
	jumlahmotor := motor / 2
	if motor%2 != 0 {
		jumlahmotor += 1
	}
	fmt.Println(jumlahmotor)
}
