package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá mundo"))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregar página de usuários"))
}

func main() {
	//HTTP é um protocolo de comunicação - base da comunicação web

	//cliente (faz requisição) - servidor (processa requisição e envia resposta)

	// request - response

	// rotas
	// URI - Identificador do Recurso
	// Método - GET, POST, PUT, DELETE
	http.HandleFunc("/home", home)
	http.HandleFunc("/usuarios", usuarios)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
