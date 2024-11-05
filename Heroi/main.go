package main

import (
	"fmt"
	"hero-api/banco"
	"hero-api/handlers"
	"log"
	"net/http"
)

func main() {
	// Conectar ao banco de dados
	db, err := banco.ConnectDB()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	// Configurar rotas
	http.HandleFunc("/heroes", handlers.HeroHandler(db))
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Rota de teste funcionando!")
	})

	// Iniciar o servidor
	fmt.Println("Servidor iniciado na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
