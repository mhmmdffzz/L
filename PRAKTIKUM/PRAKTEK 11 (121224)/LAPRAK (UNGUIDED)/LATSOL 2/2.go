package main

import "fmt"

func main() {
	var x string
	var n int
	fmt.Print("Masukkan string yang dicari: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan jumlah string: ")
	fmt.Scan(&n)
	strings := make([]string, n)
	count := 0
	firstPos := -1
	fmt.Println("Masukkan", n, "string:")
	for i := 0; i < n; i++ {
		fmt.Scan(&strings[i])
		if strings[i] == x {
			if firstPos == -1 {
				firstPos = i
			}
			count++
		}
	}
	fmt.Println("String ditemukan:", count > 0)
	fmt.Println("Posisi pertama:", firstPos+1)
	fmt.Println("Jumlah kemunculan:", count)
	fmt.Println("Ada minimal dua kemunculan:", count >= 2)
}
