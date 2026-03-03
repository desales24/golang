package main

import "fmt"

func main() {
	// Contoh penggunaan semua format verb
	var nilaiFloat float64 = 7.25
	var benar bool = true
	var karakter rune = 'A'
	var angka int = 42
	var nama string = "Golang"

	// %f : untuk float (angka desimal)
	fmt.Printf("float         : %f\n", nilaiFloat)
	// %t : untuk boolean (true/false)
	fmt.Printf("boolean       : %t\n", benar)
	// %c : untuk karakter (rune)
	fmt.Printf("karakter      : %c\n", karakter)
	// %b : untuk bilangan biner
	fmt.Printf("biner         : %b\n", angka)
	// %x : untuk heksadesimal
	fmt.Printf("heksadesimal  : %x\n", angka)
	// %T : untuk menampilkan tipe data variabel
	fmt.Printf("tipe data     : %T\n", angka)
	// %% : untuk menampilkan tanda persen (%)
	fmt.Printf("tanda persen  : %%\n")
	// %v : format serbaguna
	fmt.Printf("serbaguna     : %v\n", nama)
	// %d : hanya untuk angka/integer
	fmt.Printf("integer       : %d\n", angka)
	// %s : hanya untuk string
	fmt.Printf("string        : %s\n", nama)

}

// %f : untuk float (angka desimal)
// %t : untuk boolean (true/false)
// %c : untuk karakter (rune)
// %b : untuk bilangan biner
// %x : untuk heksadesimal
// %T : untuk menampilkan tipe data variabel
// %% : untuk menampilkan tanda persen (%)
// %v adalah format serbaguna (bisa untuk tipe data apa saja).
// %d hanya untuk angka/integer.
// %s hanya untuk string.
