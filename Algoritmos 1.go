package algoritmos_em_golang_GustavoMalago

import (
	"fmt"
)

func main() {
	var numero1 uint64
	var numero2 uint64

	fmt.Print("Informe o primeiro numero:")
	fmt.Scan(&numero1)
	fmt.Print("Informe o segundo número")
	fmt.Scan(&numero2)

	if numero1 > numero2 {

		fmt.Print("O Primeiro numero é maior")

	} else if numero1 < numero2 {

		fmt.Print("O Segundo número é maior")

	} else {

		fmt.Print("Os dois números são iguais")
	}
}
