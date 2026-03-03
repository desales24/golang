package main

import "fmt"

func main() {
	tinggi := 10
	for baris := 0; baris < tinggi; baris++ {
		// Segitiga kiri
		for kolom := 0; kolom <= baris; kolom++ {
			fmt.Print("*")
		}
		// Spasi tengah (membentuk segitiga kosong)
		spasi := (tinggi - baris - 1) * 2
		for kolom := 0; kolom < spasi; kolom++ {
			fmt.Print(" ")
		}
		// Segitiga kanan
		for kolom := 0; kolom <= baris; kolom++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
