package main

import "fmt"

func main() {
	var nilai int = 9

	switch nilai {
	case 1,2,3:
		fmt.Println("nilai kecil.")
	case 4,5,6:
		fmt.Println("nilai sedang.")
	case 7,8,9:
		fmt.Println("nilai besar.")
	default:
		fmt.Println("nilai tidak valid.")
	}
}
