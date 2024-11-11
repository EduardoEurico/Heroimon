package dataBase

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

// Funcao para conectar ao banco de dados
func ConnectDataBase() *sql.DB {
	// Data Source Name do banco de dados
	conexao := "user=docker dbname=DbParadigmas password=postgres host=localhost sslmode=disable"	dataBase, err := sql.Open("postgres", conexao)

	if err != nil {
		log.Panic("Erro ao conectar ao banco de dados: ", err)
	}

	return dataBase
}
