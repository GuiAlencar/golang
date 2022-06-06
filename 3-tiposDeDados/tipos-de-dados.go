package main

import (
	"errors"
	"fmt"
)

func main() {
	// int8, int16, int32, int64, int
	var numero int64 = 100
	fmt.Println(numero)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	//alias
	//int32 = rune
	var numero3 rune = 123456
	fmt.Println(numero3)

	//byte = uint8
	var numero4 byte = 123
	fmt.Println(numero4)

	// float32, float64

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 123456.67
	fmt.Println(numeroReal3)

	// Strings

	var str string = "texto"
	fmt.Println(str)

	str2 := "texto2"
	fmt.Println(str2)

	//não existe char no go, ele sempre converte para tablea ask
	char := 'B'
	fmt.Println(char)

	// booleano
	var booleano1 bool = true
	fmt.Println(booleano1)

	// erro
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
