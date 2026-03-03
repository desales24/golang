package main

import "fmt"

func namaSiswa(nama string, umur int) {
	fmt.Println("Nama Siswa : ", nama)
	fmt.Println("Umur Siswa : ", umur)
}

func main() {
	namaSiswa("Fransiskus", 20)
	fmt.Println()
	namaSiswa("Desales", 19)
}