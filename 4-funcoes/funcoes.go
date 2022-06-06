package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(2, 2)
	fmt.Println(soma)

	var f = func() {
		fmt.Println("Função F")
	}

	f()

	// ---------------------------------------

	var f2 = func(txt string) {
		fmt.Println(txt)
	}

	f2("Função F2")

	// ---------------------------------------

	var f3 = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f3("Função F3")
	fmt.Println(resultado)

	// ---------------------------------------

	// se não quiser retornar a soma, basta trocar seu nome por _
	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)
}
