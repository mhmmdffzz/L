package main

import "fmt"

func main() {
	var j, hasil, v1, v2 int
	fmt.Scan(&v1, &v2)
	for j = 1; j <= v2; j++ {
		hasil = hasil + v1
	}
	fmt.Print(hasil)

}
