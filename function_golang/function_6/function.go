package main

import "fmt"

func perulanganAngka(baris int) int {
	if baris == 0 {
		return 0
	}
	fmt.Print(baris)
	return perulanganAngka(baris - 1)
}

func main() {
	perulanganAngka(1)
}