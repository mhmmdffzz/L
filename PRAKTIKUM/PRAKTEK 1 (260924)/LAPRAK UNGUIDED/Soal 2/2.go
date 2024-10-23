package main

import "fmt"

func main() {
	var (
		nama1, nim1, kelas1 string
	)

	fmt.Println("Masukkan data mahasiswa pertama:")
	fmt.Print("Masukkan Nama: ")
	fmt.Scanln(&nama1)
	fmt.Print("Masukkan NIM: ")
	fmt.Scanln(&nim1)
	fmt.Print("Masukkan Kelas: ")
	fmt.Scanln(&kelas1)

	fmt.Printf("\nPerkenalkan saya adalah %s, salah satu mahasiswa Prodi S1-IF dari kelas %s dengan NIM %s.\n", nama1, kelas1, nim1)

