package database

import (
	"database/sql"
	"log"

	_ "github.com/glebarez/sqlite" // driver sqlite puro Go
)

var DB *sql.DB

func Conectar() {
	var err error
	DB, err = sql.Open("sqlite", "./tarefas.db") // arquivo do banco será criado aqui na raiz do backend
	if err != nil {
		log.Fatal(err)
	}

	// Cria a tabela se não existir
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS tarefas (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		done BOOLEAN NOT NULL CHECK (done IN (0,1))
	);
	`

	_, err = DB.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
}
