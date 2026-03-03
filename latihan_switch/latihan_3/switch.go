package main

import "fmt"

var x int 
var y int 

func operasiMatematika() {
	fmt.Print("Masukkan nilai x : ")
	fmt.Scanf("%d", &x)
	
	fmt.Print("Masukkan nilai y : ")
	fmt.Scanf("%d", &y)
} 

func perhitungaFaktorial(j int) (result int) {
	if j > 0 {
		result = j * perhitungaFaktorial(j-1)
	}else {
		result = 1
	}
	return result
}

func main() {
	for {
		var pilihan int

		fmt.Println("Masukkan Pilihan Anda")
		fmt.Println("1. Penjumlahan")
		fmt.Println("2. Pengurangan")
		fmt.Println("3. Perkalian")
		fmt.Println("4. Pembagian")
		fmt.Println("5. Faktorial")
		fmt.Println("0. Keluar")
		fmt.Print("Pilihan : ")
		fmt.Scanf("%d", &pilihan)

		if pilihan == 0 {
			fmt.Println("Program selesai.")
			break
		}

		switch pilihan {
		case 1:
			operasiMatematika()
			fmt.Println("Hasil penjumlahan", x, "+", y, "=", x+y)
		case 2:
			operasiMatematika()
			fmt.Println("Hasil pengurangan", x, "-", y, "=", x-y)
		case 3:
			operasiMatematika()
			fmt.Println("Hasil perkalian", x, "*", y, "=", x*y)
		case 4:
			operasiMatematika()
			if y != 0 {
				fmt.Println("Hasil pembagian", x, "/", y, "=", x/y)
			} else {
				fmt.Println("Error: Pembagian dengan nol tidak diperbolehkan.")
			}
		case 5:
			var j int

			fmt.Print("Factorial dari x : ")
			fmt.Scanf("%d", &j)

			for i := 1; i <= j; i++ {
				if i == j {
					fmt.Print(i, " = ", perhitungaFaktorial(i))
				} else {
					fmt.Print(i, " * ")
				}
			}
			fmt.Println()
		default:
			fmt.Println("Pilihan tidak valid.")
		}
		fmt.Println()
	}
}
