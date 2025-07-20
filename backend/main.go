package main

import (
	"log"
	"net/http"

	"todolist/routes" // importa o pacote que vai criar e configurar o mux
	"todolist/database" // importa o pacote que conecta ao banco de dados
)

func main() {
	database.Conectar() // conecta ao SQLite e cria tabela

	router := routes.HandleRequests() // chama função que cria o mux e configura as rotas

	log.Println("Servidor rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
