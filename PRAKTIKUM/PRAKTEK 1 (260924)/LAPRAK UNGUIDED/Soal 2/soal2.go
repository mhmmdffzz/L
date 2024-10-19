package main

import "fmt"

func main() {
	var (
		nama1, nim1, kelas1 string
		nama2, nim2, kelas2 string
	)

	// Input data mahasiswa pertama
	fmt.Println("Masukkan data mahasiswa pertama:")
	fmt.Print("Masukkan Nama: ")
	fmt.Scanln(&nama1)
	fmt.Print("Masukkan NIM: ")
	fmt.Scanln(&nim1)
	fmt.Print("Masukkan Kelas: ")
	fmt.Scanln(&kelas1)

	// Input data mahasiswa kedua
	fmt.Println("\nMasukkan data mahasiswa kedua:")
	fmt.Print("Masukkan Nama: ")
	fmt.Scanln(&nama2)
	fmt.Print("Masukkan NIM: ")
	fmt.Scanln(&nim2)
	fmt.Print("Masukkan Kelas: ")
	fmt.Scanln(&kelas2)

	// Menampilkan biodata mahasiswa pertama
	fmt.Printf("\nPerkenalkan saya adalah %s, salah satu mahasiswa Prodi S1-IF dari kelas %s dengan NIM %s.\n", nama1, kelas1, nim1)

	// Menampilkan biodata mahasiswa kedua
	fmt.Printf("Perkenalkan saya adalah %s, salah satu mahasiswa Prodi S1-IF dari kelas %s dengan NIM %s.\n", nama2, kelas2, nim2)
}
