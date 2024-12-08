package main

import "fmt"

func main() {
	var target, donasi, totaldonasi, jumlahdonatur int
	fmt.Scan(&target)
	totaldonasi = 0
	jumlahdonatur = 0
	for {
		fmt.Scan(&donasi)
		totaldonasi += donasi
		jumlahdonatur++
		fmt.Printf("jumlahdonatur %d : menyumbang %d."+"total terkumpul : %d\n", jumlahdonatur, donasi, totaldonasi)
		if totaldonasi >= target {
			break
		}
	}
	fmt.Printf("Target tercapai! Total donasi:"+
		"%d dari %d donatur.", totaldonasi, jumlahdonatur)
}
