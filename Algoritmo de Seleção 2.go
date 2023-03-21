package main

import "fmt"

func main() {
	var x int64
	var y int64
	var z int64
	var menor int64

	fmt.Println("Informe o primeiro numero: ")
	fmt.Scan(&x)
	fmt.Println("Informe o segundo numero: ")
	fmt.Scan(&y)
	fmt.Println("Infome o terceiro numero: ")
	fmt.Scan(&z)

	if x < y && x < z {
		menor = x
	} else if y < x && y < z {
		menor = y
	} else {
		menor = z
	}
	{
		fmt.Print("O menor número é:", menor)
	}
}