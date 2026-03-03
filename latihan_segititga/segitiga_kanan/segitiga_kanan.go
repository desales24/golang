package main

import "fmt"

func main() {
	// Loop untuk setiap baris
	for baris := 0; baris < 10; baris++ {
		// Loop untuk mencetak spasi di sebelah kiri
		for kolom := 0; kolom < 10-baris-1; kolom++ {
			fmt.Print(" ")
		}
		// Loop untuk mencetak bintang di sebelah kanan
		for kolom := 0; kolom < baris+1; kolom++ {
			fmt.Print("*")
		}
		// Pindah ke baris berikutnya
		fmt.Println()
	}
}
