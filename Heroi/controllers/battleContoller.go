// controllers/battle_controller.go
package controllers

import (
	"hero-api/dataBase"
	models2 "hero-api/models"
	"hero-api/view"

	"log"
	"net/http"
	"strconv"
)

// ListBattles displays all battle records.
func ListBattles(w http.ResponseWriter, r *http.Request) {
	db := dataBase.ConnectDataBase()
	defer db.Close()

	battles, err := models2.GetAllBattleRecords(db)
	if err != nil {
		log.Println("Erro ao obter registros de batalha:", err)
		http.Error(w, "Erro ao obter registros de batalha", http.StatusInternalServerError)
		return
	}

	view.Templates.ExecuteTemplate(w, "ListBattles", battles)
}

// StartBattle renders the page to start a new battle.
func StartBattle(w http.ResponseWriter, r *http.Request) {
	db := dataBase.ConnectDataBase()
	defer db.Close()

	// Get all heroes to select for the battle
	heroes, err := models2.GetAllHeroes(db)
	if err != nil {
		log.Println("Erro ao obter heróis:", err)
		http.Error(w, "Erro ao obter heróis", http.StatusInternalServerError)
		return
	}

	view.Templates.ExecuteTemplate(w, "StartBattle", heroes)
}

// BattleResult handles the battle simulation and displays the result.
func BattleResult(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Parse form data
		err := r.ParseForm()
		if err != nil {
			log.Println("Erro ao parsear o formulário:", err)
			http.Error(w, "Erro ao processar o formulário", http.StatusBadRequest)
			return
		}

		// Extract form values
		heroi1ID, _ := strconv.Atoi(r.FormValue("heroi1_id"))
		heroi2ID, _ := strconv.Atoi(r.FormValue("heroi2_id"))

		// Simulate battle
		db := dataBase.ConnectDataBase()
		defer db.Close()

		resultado, err := models2.StartBattle(db, heroi1ID, heroi2ID)
		if err != nil {
			log.Println("Erro ao simular batalha:", err)
			http.Error(w, "Erro ao simular batalha", http.StatusInternalServerError)
			return
		}

		// Display the result
		view.Templates.ExecuteTemplate(w, "BattleResult", resultado)
	}
}

// BattleDetails displays details of a specific battle.
func BattleDetails(w http.ResponseWriter, r *http.Request) {
	// Get the battle record ID from the URL query
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("ID inválido:", err)
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	db := dataBase.ConnectDataBase()
	defer db.Close()

	battle, err := models2.GetBattleRecordByID(db, id)
	if err != nil {
		log.Println("Erro ao obter registro de batalha:", err)
		http.Error(w, "Erro ao obter registro de batalha", http.StatusInternalServerError)
		return
	}

	view.Templates.ExecuteTemplate(w, "BattleDetails", battle)
}
