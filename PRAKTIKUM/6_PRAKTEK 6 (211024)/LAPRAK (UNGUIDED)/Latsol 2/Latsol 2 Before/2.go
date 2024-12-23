package main

import "fmt"

func main() {
	var nam float64
	var nmk string
	fmt.Print("Nilai akhir mata kuliah: ")
	fmt.Scan(&nam)
	if nam > 80 {
		nmk = "A"
	}
	if nam > 72.5 {
		nmk = "AB"
	}
	if nam > 65 {
		nmk = "B"
	}
	if nam > 57.5 {
		nmk = "BC"
	}
	if nam > 50 {
		nmk = "C"
	}
	if nam > 40 {
		nmk = "D"
	} else if nam <= 40 {
		nmk = "E"
	}
	fmt.Println("Nilai mata kuliah: ", nmk)
}
