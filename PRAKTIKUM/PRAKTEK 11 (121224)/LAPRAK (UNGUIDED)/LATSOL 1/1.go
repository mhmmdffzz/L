package main

import "fmt"

func main() {
	var num float64
	sum := 0.0
	count := 0
	fmt.Println("Masukkan bilangan (9999 untuk berhenti):")
	for {
		fmt.Scan(&num)
		if num == 9999 {
			break
		}
		sum += num
		count++
	}
	if count > 0 {
		average := sum / float64(count)
		fmt.Printf("Rata-rata: %.2f\n", average)
	} else {
		fmt.Println("Tidak ada bilangan yang dimasukkan")
	}
}
