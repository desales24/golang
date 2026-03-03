package main

import "fmt"

func namaSiswa(nama_depan string, nama_belakang string) (result string) {
	result = nama_depan + " " + nama_belakang                 
	return result
}

func main() {
	fmt.Println("Nama Siswa:", namaSiswa("Fransiskus", "Desales"))
}
