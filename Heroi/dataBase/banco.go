package dataBase

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// Funcao para conectar ao banco de dados
func ConnectDataBase() *sql.DB {
	// Data Source Name do banco de dados
	conexao := "user=postgres dbname=postgres password=postgres host=localhost port=5433 sslmode=disable"

	dataBase, err := sql.Open("postgres", conexao)

	if err != nil {
		log.Panic("Erro ao conectar ao banco de dados: ", err)
	}

	return dataBase
}
