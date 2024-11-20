package main

import "fmt"

func main() {
	var bilangan int
	fmt.Scan(&bilangan)
	if bilangan < 0 {
		bilangan = -bilangan
	}
	fmt.Print(bilangan)

}
