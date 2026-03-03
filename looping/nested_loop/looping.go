package main

import "fmt"

func main () {
	baris := [5] int {1,2,3,4,5}
	kolom := [5] int {6,7,8,9,10}

	for i := 0; i < len(baris); i++ {
		for j := 0; j < len(kolom); j++ {
			fmt.Print(baris[i], kolom[j], " ")
		}
		fmt.Println()
	}
}
