package main

import "fmt"

type date struct {
	tanggal int
	bulan   string
	tahun   int
}

func main() {
	var tgl date
	tgl.tanggal = 10
	fmt.Scan(&tgl.bulan, &tgl.tahun)
	tgl.tahun += 10
	fmt.Println(tgl.tanggal)
	fmt.Println(tgl.bulan, tgl.tahun)
}
