package routes

import (
	"github.com/gorilla/mux"
	"todolist/controllers" // ajuste conforme o nome do seu módulo
)

// HandleRequests cria o roteador mux, registra as rotas e retorna ele
func HandleRequests() *mux.Router {
	router := mux.NewRouter()

	// Rotas para a API de tarefas
	router.HandleFunc("/tarefas", controllers.ObterTarefas).Methods("GET")
	router.HandleFunc("/tarefas", controllers.CriarTarefa).Methods("POST")
	router.HandleFunc("/tarefas/{id}", controllers.ExcluirTarefa).Methods("DELETE")

	return router
}