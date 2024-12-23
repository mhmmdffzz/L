package main

import "fmt"

func main() {
	var x rune
	var huruf, vKecil, vBesar bool
	fmt.Scanf("%c", &x)
	huruf = (x >= 'a' && x <= 'z' || x >= 'A' && x <= 'Z')
	vKecil = x == 'a' || x == 'i' || x == 'u' || x == 'e' || x == 'o'
	vBesar = x == 'A' || x == 'I' || x == 'U' || x == 'E' || x == 'O'
	if huruf && (vKecil || vBesar) {
		fmt.Println("vokal")
	} else if huruf && !(vKecil || vBesar) {
		fmt.Println("konsonan")
	} else {
		fmt.Println("bukan huruf")
	}

}
