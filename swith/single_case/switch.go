package main

import "fmt"

func main () {
	var nilai int = 80

	switch {
		case nilai >= 80:
			fmt.Println ("kamu mendapat nilai A.")
		case nilai >= 70:
			fmt.Println ("kamu mendapat nilai B.")
		case nilai >= 60:
			fmt.Println ("kamu mendapat nilai C.")
		default:
			fmt.Println ("melihat nilai kamu.")
	} 
}
