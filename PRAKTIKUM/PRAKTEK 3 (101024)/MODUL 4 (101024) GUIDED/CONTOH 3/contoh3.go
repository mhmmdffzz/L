package main

import "fmt"

func main() {
	var beratbadan, tinggibadan, bmi float64
	fmt.Print("Masukkan berat badan (kg) : ")
	fmt.Scanln(&beratbadan)
	fmt.Print("Masukkan tinggi badan (m) : ")
	fmt.Scanln(&tinggibadan)
	bmi = beratbadan / (tinggibadan * tinggibadan)
	fmt.Printf("BMI anda: %.2f", bmi)
}
