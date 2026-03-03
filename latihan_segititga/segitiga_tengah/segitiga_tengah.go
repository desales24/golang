package main

import "fmt"

func main() {
	tinggi := 12
	for baris := 0; baris < tinggi; baris++ {
		// Cetak spasi di sebelah kiri
		for kolom := 0; kolom < tinggi-baris-1; kolom++ {
			fmt.Print(" ")
		}
		// Cetak bintang
		for kolom := 0; kolom < (baris+1)*2-1; kolom++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
