package main

import "fmt"

func perhitungaFactorial(x int) (result int) {
	if x > 0 {
		result = x * perhitungaFactorial(x-1)
	}else {
		result = 1
	}
	return result
}

func main() {
	var x int

	fmt.Print("Factorial dari x : ")
	fmt.Scanf("%d", &x)

	for i := 1; i <= x; i++ {
		if i == x {
			fmt.Print(i, " = ", perhitungaFactorial(i))
		} else {
			fmt.Print(i, " * ")
		}
	}
}