package main 

import "fmt"

func main () {
	baris := 10

	for baris=0 ; baris<10; baris ++ {
		if baris == 5 {
			continue
		}
		fmt.Println(baris)
	}
}