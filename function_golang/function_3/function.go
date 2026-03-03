package main

import "fmt"

func namaSiswa(nama_depan string, nama_belakang string) string {
	return nama_depan + " " + nama_belakang                 
}

func main() {
	fmt.Println(namaSiswa("Fransiskus", "Desales"))
}
