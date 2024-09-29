package main

import "fmt"

func main() {
	var fahrenheit float64

	fmt.Print("Masukan suhu dalam Fahrenheit: ")
	fmt.Scanln(&fahrenheit)

	celcius := (fahrenheit - 32) * 5 / 9

	fmt.Printf("Suhu %.2f Fahrenheit adalah %.2f Celcius\n", fahrenheit, celcius)
}
