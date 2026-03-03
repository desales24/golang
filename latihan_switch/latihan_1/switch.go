package main

import "fmt"

func main() {
	var x int
	var y int

	fmt.Print("Masukkan nilai x : ")
	fmt.Scanf("%d", &x)

	fmt.Print("Masukkan nilai y : ")
	fmt.Scanf("%d", &y)

	fmt.Println("Hasil penjumlahan", x, "+", y, "=", x+y)
}
