package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalComPonteiros(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	//---------------------------------------------

	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiros(&novoNumero)
	fmt.Println(novoNumero)
}
