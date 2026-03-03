package main

import "fmt"

func main () {
	// Loop untuk setiap baris
	for baris := 0; baris < 10; baris++ {
		// Loop untuk mencetak bintang sebanyak nomor baris+1
		for kolom := 0; kolom < baris+1; kolom++ {
			fmt.Print("*")
		}
		// Pindah ke baris berikutnya
		fmt.Println()
	}
}
