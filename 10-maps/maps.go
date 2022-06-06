package main

import "fmt"

func main() {

	usuario := map[string]string{
		"nome":      "Guilherme",
		"sobrenome": "Alencar",
	}
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Guilherme",
			"ultimo":   "Alencar",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "campus 1",
		},
	}
	fmt.Println(usuario2)
	// apagando chave do map
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	//adicionando uma chave no map
	usuario2["signo"] = map[string]string{
		"nome": "Sagitário",
	}
	fmt.Println(usuario2["signo"])
}
