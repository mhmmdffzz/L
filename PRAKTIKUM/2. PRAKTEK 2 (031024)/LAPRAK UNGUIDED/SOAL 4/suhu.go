package main

import "fmt"

func main() {
	var fahrenheit, celcius, reamur, kelvin int
	//  Input suhu
	fmt.Print("Temperatur Celcius: ")
	fmt.Scan(&celcius)
	//  Konversi celcius ke fahrenheit
	fahrenheit = int((float64(celcius) * 9 / 5) + 32)
	reamur = int(float64(celcius) * 4 / 5)
	kelvin = int(float64(celcius) + 273.15)
	//   Output
	fmt.Println("Derajat Fahrenheit: ", fahrenheit)
	fmt.Println("Derajat Reamur:  ", reamur)
	fmt.Println("Derajat kelvin:  ", kelvin)

}
