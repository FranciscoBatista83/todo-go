#!/bin/bash

echo "íº§ Criando estrutura do backend..."

# CriaÃ§Ã£o das pastas
mkdir -p backend/{models,controllers,routes}

# CriaÃ§Ã£o dos arquivos
touch backend/main.go
touch backend/models/task.go
touch backend/controllers/task_controller.go
touch backend/routes/task_routes.go

# ConteÃºdo do main.go
cat > backend/main.go <<EOF
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"todo-list/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterTaskRoutes(r)

	log.Println("Servidor rodando na porta 8080 íº€")
	log.Fatal(http.ListenAndServe(":8080", r))
}
EOF

# ConteÃºdo do models/task.go
cat > backend/models/task.go <<EOF
package models

type Task struct {
	ID    string \`json:"id"\`
	Title string \`json:"title"\`
	Done  bool   \`json:"done"\`
}
EOF

# ConteÃºdo do controllers/task_controller.go
cat > backend/controllers/task_controller.go <<EOF
package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"todo-list/models"
)

var tasks []models.Task

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var task models.Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	tasks = append(tasks, task)
	json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, task := range tasks {
		if task.ID == params["id"] {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}
	w.WriteHeader(http.StatusNoContent)
}
EOF

# ConteÃºdo do routes/task_routes.go
cat > backend/routes/task_routes.go <<EOF
package routes

import (
	"github.com/gorilla/mux"
	"todo-list/controllers"
)

func RegisterTaskRoutes(router *mux.Router) {
	router.HandleFunc("/tasks", controllers.GetTasks).Methods("GET")
	router.HandleFunc("/tasks", controllers.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", controllers.DeleteTask).Methods("DELETE")
}
EOF

echo "âœ… Estrutura criada com sucesso!"
