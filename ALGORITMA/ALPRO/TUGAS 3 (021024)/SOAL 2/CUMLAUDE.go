package main

import "fmt"

func main() {
	var semester, eprt int
	var cumlaude bool

	fmt.Print("Masukkan jumlah semester dan skor EPrT: ")
	fmt.Scan(&semester, &eprt)

	cumlaude = (semester <= 8) && (eprt >= 500)
	fmt.Println(cumlaude)
}
