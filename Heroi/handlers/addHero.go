// addHero.go
package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

// Estrutura do herói
type Hero struct {
	NomeReal        string  `json:"nome_real"`
	NomeHeroi       string  `json:"nome_heroi"`
	Sexo            string  `json:"sexo"`
	Altura          float64 `json:"altura"`
	Peso            float64 `json:"peso"`
	DataNascimento  string  `json:"data_nascimento"`
	LocalNascimento string  `json:"local_nascimento"`
}

// HeroHandler gerencia a adição de novos heróis
func HeroHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}

		var hero Hero
		err := json.NewDecoder(r.Body).Decode(&hero)
		if err != nil {
			http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
			return
		}

		err = insertHero(db, hero)
		if err != nil {
			http.Error(w, "Erro ao inserir herói no banco de dados", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "Herói %s adicionado com sucesso!", hero.NomeHeroi)
	}
}

// Função para inserir o herói no banco
func insertHero(db *sql.DB, hero Hero) error {
	query := `
		INSERT INTO heroes (nome_real, nome_heroi, sexo, altura, peso, data_nascimento, local_nascimento)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`
	_, err := db.Exec(query, hero.NomeReal, hero.NomeHeroi, hero.Sexo, hero.Altura, hero.Peso, hero.DataNascimento, hero.LocalNascimento)
	return err
}
