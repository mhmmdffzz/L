package main

import "fmt"

func main() {
	var sisi, volume float64
	fmt.Scan(&sisi)
	volume = (sisi * sisi * sisi)
	fmt.Print(volume)
}
