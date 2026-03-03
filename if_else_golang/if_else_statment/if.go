package main

import "fmt"

func main() {
	var nilai int = 75

	if nilai >= 80 {
		fmt.Println ("nilai kamu bagus")
	}else if nilai < 80 && nilai >= 60 {
		fmt.Println("nilai kamu kurang bagus")
	}else {
		fmt.Println("nilai kamu buruk")
	}
}