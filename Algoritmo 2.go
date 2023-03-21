package main

import "fmt"

func main() {
	var x uint64

	fmt.Print("Informe o numero que deseje:")
	fmt.Scan(&x)
	fmt.Println("O dobro desse numero é igual a ", x*x)
	fmt.Println("O triplo desse numeoro é igual a ", x*x*x)
	fmt.Println("O quadruplo desse numero é igual a ", x*x*x*x)
}
