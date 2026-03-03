package main

import "fmt"

func main () {
	var (
		nama string = "uyab"
		umur int = 20 
		jenis_kelamin string = "laki-laki"
		a int = 10 
		b int = 10
		penjumlahan int = a + b
	)

	harusdalamfunc := "harus dalam func main"

	fmt.Println(nama)
	fmt.Println(umur)
	fmt.Println(jenis_kelamin)
	fmt.Println(penjumlahan)
	fmt.Println(harusdalamfunc)
}