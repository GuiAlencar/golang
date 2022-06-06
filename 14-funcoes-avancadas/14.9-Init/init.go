package main

import "fmt"

// função init é sempre inicializada primeiro que o main
func init() {
	fmt.Println("Executando a função init")
}
func main() {
	fmt.Println("Executando a função main")
}
