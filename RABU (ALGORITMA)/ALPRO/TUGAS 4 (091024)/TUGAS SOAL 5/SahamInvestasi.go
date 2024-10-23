package main

import "fmt"

func hitungKeuntunganBersih(hargaBeli, hargaJual, jumlahSaham float64) float64 {
	totalInvestasiAwal := hargaBeli * jumlahSaham
	totalPenjualan := hargaJual * jumlahSaham
	keuntunganKotor := totalPenjualan - totalInvestasiAwal
	biayaTransaksi := 0.002 * totalPenjualan
	var pajakKeuntungan float64
	if keuntunganKotor > 0 {
		pajakKeuntungan = 0.10 * keuntunganKotor
	} else {
		pajakKeuntungan = 0
	}
	keuntunganBersih := keuntunganKotor - biayaTransaksi - pajakKeuntungan
	return keuntunganBersih
}

func main() {
	var hargaBeli, hargaJual, jumlahSaham float64
	fmt.Scan(&hargaBeli)
	fmt.Scan(&hargaJual)
	fmt.Scan(&jumlahSaham)
	keuntunganBersih := hitungKeuntunganBersih(hargaBeli, hargaJual, jumlahSaham)
	fmt.Printf("Keuntungan bersih dari investasi: %.2f\n", keuntunganBersih)
}
