package main

import "fmt"

func main() {
	for {
		var x int
		var y int
		var pilihan int

		fmt.Print("Masukkan nilai x : ")
		fmt.Scanf("%d", &x)

		fmt.Print("Masukkan nilai y : ")
		fmt.Scanf("%d", &y)

		fmt.Println("Masukkan Pilihan Anda")
		fmt.Println("1. Penjumlahan")
		fmt.Println("2. Pengurangan")
		fmt.Println("3. Perkalian")
		fmt.Println("4. Pembagian")
		fmt.Println("5. Backup")
		fmt.Println("0. Keluar")
		fmt.Print("Pilihan : ")
		fmt.Scanf("%d", &pilihan)

		if pilihan == 0 {
			fmt.Println("Program selesai.")
			break
		}

		switch pilihan {
		case 1:
			fmt.Println("Hasil penjumlahan", x, "+", y, "=", x+y)
		case 2:
			fmt.Println("Hasil pengurangan", x, "-", y, "=", x-y)
		case 3:
			fmt.Println("Hasil perkalian", x, "*", y, "=", x*y)
		case 4:
			if y != 0 {
				fmt.Println("Hasil pembagian", x, "/", y, "=", x/y)
			} else {
				fmt.Println("Error: Pembagian dengan nol tidak diperbolehkan.")
			}
		case 5:
			fmt.Println("Backup dipilih. Tidak ada operasi matematika yang dilakukan.")
		default:
			fmt.Println("Pilihan tidak valid.")
		}
		fmt.Println()
	}
}
