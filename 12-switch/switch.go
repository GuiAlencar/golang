package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "quarta"
	case 5:
		return "quinta"
	case 6:
		return "sexta"
	case 7:
		return "Sábado"
	default:
		return "Número inválido"
	}
}

func diaDaSemana2(numero int) string {
	var diaDaSemana string
	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		//fallthrough faz cair na próxima condição numero 1 cairia na segunda-feira
	case numero == 2:
		diaDaSemana = "Segunda"
	case numero == 3:
		diaDaSemana = "terça"
	case numero == 4:
		diaDaSemana = "Quarta"
	case numero == 5:
		diaDaSemana = "quinta"
	case numero == 6:
		diaDaSemana = "sexta"
	case numero == 7:
		diaDaSemana = "Sãbado"
	default:
		diaDaSemana = "Número inválido"
	}
	return diaDaSemana
}

func main() {

	dia := diaDaSemana(10)
	fmt.Println(dia)

	fmt.Println("------------------------------")
	dia2 := diaDaSemana2(3)
	fmt.Println(dia2)
}
