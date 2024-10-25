package banco

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// ConnectDB faz a conexão com o banco de dados e retorna o *sql.DB
func ConnectDB() (*sql.DB, error) {
	dsn := "root:ceub123456@tcp(127.0.0.1:3306)/heroes_db" // Substitua username e password
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Conectado ao banco de dados com sucesso!")
	return db, nil
}
