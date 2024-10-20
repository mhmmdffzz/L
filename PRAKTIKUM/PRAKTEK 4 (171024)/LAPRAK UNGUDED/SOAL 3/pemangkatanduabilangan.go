package main

import "fmt"

func main() {
	var bilanganbulat1, bilanganbulat2 int
	fmt.Scan(&bilanganbulat1, &bilanganbulat2)
	hasil := 1
	for i := 0; i < bilanganbulat2; i++ {
		hasil *= bilanganbulat1
	}
	fmt.Println(hasil)
}
