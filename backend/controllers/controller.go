package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"todolist/database"
	"todolist/models"
)

// Obter todas as tarefas do banco
func ObterTarefas(w http.ResponseWriter, r *http.Request) {
	// Define o tipo de conteúdo da resposta como JSON
	w.Header().Set("Content-Type", "application/json")

	// Executa a query para buscar todas as tarefas
	rows, err := database.DB.Query("SELECT id, title, done FROM tarefas")
	if err != nil {
		// Em caso de erro, retorna status 500 (erro interno)
		http.Error(w, "Erro ao consultar banco", http.StatusInternalServerError)
		return
	}
	// Garante que o resultado será fechado ao final da função
	defer rows.Close()

	// Cria um slice para armazenar as tarefas buscadas
	var tarefas []models.Tarefa

	// Itera sobre as linhas retornadas pela query
	for rows.Next() {
		var t models.Tarefa
		var doneInt int
		// Faz o scan dos dados para as variáveis da struct e doneInt
		err := rows.Scan(&t.ID, &t.Title, &doneInt)
		if err != nil {
			// Se der erro lendo dados, retorna status 500
			http.Error(w, "Erro ao ler dados", http.StatusInternalServerError)
			return
		}
		// Converte o valor do campo done (int) para bool
		t.Done = doneInt == 1
		// Adiciona a tarefa ao slice
		tarefas = append(tarefas, t)
	}

	// Codifica e envia o slice de tarefas como JSON na resposta
	json.NewEncoder(w).Encode(tarefas)
}

// Criar nova tarefa no banco
func CriarTarefa(w http.ResponseWriter, r *http.Request) {
	// Define o tipo de conteúdo da resposta como JSON
	w.Header().Set("Content-Type", "application/json")

	// Decodifica o JSON enviado na requisição para a struct tarefa
	var tarefa models.Tarefa
	if err := json.NewDecoder(r.Body).Decode(&tarefa); err != nil {
		// Se JSON inválido, retorna erro 400 (bad request)
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	// Converte bool para int para armazenar no banco (0 ou 1)
	doneInt := 0
	if tarefa.Done {
		doneInt = 1
	}

	// Executa o comando para inserir uma nova tarefa no banco
	result, err := database.DB.Exec("INSERT INTO tarefas (title, done) VALUES (?, ?)", tarefa.Title, doneInt)
	if err != nil {
		// Se erro na inserção, retorna erro 500
		http.Error(w, "Erro ao inserir tarefa", http.StatusInternalServerError)
		return
	}

	// Obtém o ID gerado automaticamente pelo banco para a nova tarefa
	id, err := result.LastInsertId()
	if err != nil {
		// Se erro ao obter ID, retorna erro 500
		http.Error(w, "Erro ao obter ID da tarefa", http.StatusInternalServerError)
		return
	}

	// Atualiza o campo ID da struct tarefa com o ID do banco
	tarefa.ID = int(id)

	// Retorna status 201 (created) e envia a tarefa criada como JSON
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tarefa)
}

// Excluir tarefa pelo ID no banco
func ExcluirTarefa(w http.ResponseWriter, r *http.Request) {
	// Define o tipo de conteúdo da resposta como JSON
	w.Header().Set("Content-Type", "application/json")

	// Obtém os parâmetros da URL (no caso, o ID da tarefa)
	params := mux.Vars(r)
	idParam := params["id"]

	// Converte o ID recebido (string) para int
	id, err := strconv.Atoi(idParam)
	if err != nil {
		// Se o ID não for um número válido, retorna erro 400
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	// Executa o comando para excluir a tarefa com o ID informado
	result, err := database.DB.Exec("DELETE FROM tarefas WHERE id = ?", id)
	if err != nil {
		// Se erro ao excluir, retorna erro 500
		http.Error(w, "Erro ao excluir tarefa", http.StatusInternalServerError)
		return
	}

	// Verifica quantas linhas foram afetadas pela exclusão
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		// Se erro ao verificar, retorna erro 500
		http.Error(w, "Erro ao verificar exclusão", http.StatusInternalServerError)
		return
	}

	// Se nenhuma linha foi afetada, significa que a tarefa não foi encontrada
	if rowsAffected == 0 {
		http.Error(w, "Tarefa não encontrada", http.StatusNotFound)
		return
	}

	// Retorna status 204 (No Content) para indicar que exclusão foi bem-sucedida
	w.WriteHeader(http.StatusNoContent)
}
