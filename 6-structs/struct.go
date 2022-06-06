package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo Structs")
	var u usuario
	u.nome = "Guilherme"
	u.idade = 24
	fmt.Println(u)
	fmt.Println(u.nome)
	fmt.Println(u.idade)

	fmt.Println("------------------------------------------")

	enderecoExemplo := endereco{"Rua tal", 60}

	usuario2 := usuario{"Guilherme", 24, enderecoExemplo}
	fmt.Println(usuario2)
	fmt.Println(usuario2.nome)
	fmt.Println(usuario2.idade)
	fmt.Println(usuario2.endereco.logradouro)

	fmt.Println("------------------------------------------")
	// Passando apenas um valor para um struct com dois valores
	usuario3 := usuario{nome: "Guilherme"}
	fmt.Println(usuario3)
	fmt.Println(usuario3.nome)
	fmt.Println(usuario3.endereco.numero)

}
