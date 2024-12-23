package main

import "fmt"

func main() {
	var detik, jam, menit int
	fmt.Scan(&detik)
	jam = detik / 3600
	menit = (detik % 3600) / 60
	detik = detik % 60
	fmt.Print(jam, "jam", menit, "menit", detik, "detik")
}
