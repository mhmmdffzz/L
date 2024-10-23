package main

import (
	"fmt"
)

type Transaksi struct {
	NamaBarang  string
	Jumlah      int
	HargaSatuan float64
}

func (t Transaksi) HitungTotalHarga() float64 {
	return float64(t.Jumlah) * t.HargaSatuan
}
func main() {
	namaBarang := "Pensil"
	jumlah := 20
	hargaSatuan := 2000.0

	transaksi := Transaksi{
		NamaBarang:  namaBarang,
		Jumlah:      jumlah,
		HargaSatuan: hargaSatuan,
	}
	totalHarga := transaksi.HitungTotalHarga()
	fmt.Printf("Informasi Transaksi:\n")
	fmt.Printf("Nama Barang: %s\n", transaksi.NamaBarang)
	fmt.Printf("Jumlah: %d\n", transaksi.Jumlah)
	fmt.Printf("Harga Satuan: Rp %.f\n", transaksi.HargaSatuan)
	fmt.Printf("Total Harga: Rp %.2f\n", totalHarga)
}
