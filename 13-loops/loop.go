package main

import (
	"fmt"
)

func main() {
	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println(i)
	// }

	// fmt.Println("----------------------")

	// for j := 0; j < 10; j += 2 {
	// 	fmt.Println("Incrementado J", j)
	// 	time.Sleep(time.Second)
	// }

	// fmt.Println("----------------------")

	nomes := [3]string{"João", "Davi", "Lucas"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Guilherme",
		"sobrenome": "Alencar",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}
