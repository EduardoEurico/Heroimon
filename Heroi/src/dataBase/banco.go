package dataBase

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hero-api/src/models"
	"log"
)

var (
	DB  *gorm.DB
	err error
)

// Funcao para conectar ao banco de dados
func ConnectDataBase() {
	// Data Source Name do banco de dados
	dataSourceName := "root:ceub123456@tcp(127.0.0.1:3306)/heroes_db" // Atualize com as credenciais corretas
	DB, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar com o banco de dados", err)
	}

	fmt.Println("Conectado ao banco de dados com sucesso!")
	DB.AutoMigrate(&models.Hero{})
}
