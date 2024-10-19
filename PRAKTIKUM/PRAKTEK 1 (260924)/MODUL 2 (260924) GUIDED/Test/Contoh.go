package main

import "fmt"

func main() {
	var num1, num2, num3, num4, num5, num6, num7, num8, num9, num10, num11, num12, num13, num14 int

	// Mengambil input angka untuk membentuk nama "Muhammad Fauzan"
	fmt.Print("Masukkan 14 angka untuk nama Muhammad Fauzan (gunakan angka sesuai dengan ASCII): ")
	fmt.Scan(&num1, &num2, &num3, &num4, &num5, &num6, &num7, &num8, &num9, &num10, &num11, &num12, &num13, &num14)

	// Mencetak hasil karakter berdasarkan angka yang dimasukkan
	fmt.Printf("Karakter yang dibentuk: %c%c%c%c%c%c%c%c%c%c%c%c%c%c\n", 
		byte(num1), byte(num2), byte(num3), byte(num4), byte(num5), 
		byte(num6), byte(num7), byte(num8), byte(num9), byte(num10), 
		byte(num11), byte(num12), byte(num13), byte(num14))
}
