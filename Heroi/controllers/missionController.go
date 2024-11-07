package controllers

import (
	"hero-api/dataBase"
	models2 "hero-api/models"
	"hero-api/view"
	"log"
	"net/http"
	"strconv"
)

// ListMissions displays all missions.
func ListMissions(w http.ResponseWriter, r *http.Request) {
	db := dataBase.ConnectDataBase()
	defer db.Close()

	missions, err := models2.GetAllMissions(db)
	if err != nil {
		log.Println("Erro ao obter missões:", err)
		http.Error(w, "Erro ao obter missões", http.StatusInternalServerError)
		return
	}

	view.Templates.ExecuteTemplate(w, "ListMissions", missions)
}

// CreateMission renders the page to create a new mission.
func CreateMission(w http.ResponseWriter, r *http.Request) {
	db := dataBase.ConnectDataBase()
	defer db.Close()

	// Get all heroes to assign to the mission
	heroes, err := models2.GetAllHeroes(db)
	if err != nil {
		log.Println("Erro ao obter heróis:", err)
		http.Error(w, "Erro ao obter heróis", http.StatusInternalServerError)
		return
	}

	view.Templates.ExecuteTemplate(w, "CreateMission", heroes)
}

// InsertMission handles the insertion of a new mission.
func InsertMission(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Parse form data
		err := r.ParseForm()
		if err != nil {
			log.Println("Erro ao parsear o formulário:", err)
			http.Error(w, "Erro ao processar o formulário", http.StatusBadRequest)
			return
		}

		// Extract form values
		nomeMissao := r.FormValue("nome_missao")
		descricao := r.FormValue("descricao")
		nivelDificuldade, _ := strconv.Atoi(r.FormValue("nivel_dificuldade"))
		recompensa, _ := strconv.Atoi(r.FormValue("recompensa"))
		resultado := r.FormValue("resultado")

		// Get selected heroes (multiple selection)
		heroisDesignadosStr := r.Form["herois_designados"]
		var heroisDesignados []int
		for _, idStr := range heroisDesignadosStr {
			id, _ := strconv.Atoi(idStr)
			heroisDesignados = append(heroisDesignados, id)
		}

		// Create a new Mission object
		newMission := models2.Mission{
			MissionName:  nomeMissao,
			Description:  descricao,
			Dificulty:    nivelDificuldade,
			HeroAssigned: heroisDesignados,
			Result:       resultado,
			Reward:       recompensa,
		}

		// Insert into database
		db := dataBase.ConnectDataBase()
		defer db.Close()

		err = models2.AddMission(db, newMission)
		if err != nil {
			log.Println("Erro ao adicionar missão:", err)
			http.Error(w, "Erro ao adicionar missão", http.StatusInternalServerError)
			return
		}

		// Redirect to the list of missions
		http.Redirect(w, r, "/missions", http.StatusSeeOther)
	}
}

// DeleteMission handles the deletion of a mission.
func DeleteMission(w http.ResponseWriter, r *http.Request) {
	// Get the mission ID from the URL query
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("ID inválido:", err)
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	db := dataBase.ConnectDataBase()
	defer db.Close()

	err = models2.RemoveMission(db, id)
	if err != nil {
		log.Println("Erro ao remover missão:", err)
		http.Error(w, "Erro ao remover missão", http.StatusInternalServerError)
		return
	}

	// Redirect to the list of missions
	http.Redirect(w, r, "/missions", http.StatusSeeOther)
}

// EditMission renders the page to edit an existing mission.
func EditMission(w http.ResponseWriter, r *http.Request) {
	// Get the mission ID from the URL query
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("ID inválido:", err)
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	db := dataBase.ConnectDataBase()
	defer db.Close()

	mission, err := models2.GetMissionByID(db, id)
	if err != nil {
		log.Println("Erro ao obter missão:", err)
		http.Error(w, "Erro ao obter missão", http.StatusInternalServerError)
		return
	}

	// Get all heroes to assign to the mission
	heroes, err := models2.GetAllHeroes(db)
	if err != nil {
		log.Println("Erro ao obter heróis:", err)
		http.Error(w, "Erro ao obter heróis", http.StatusInternalServerError)
		return
	}

	// Pass both mission and heroes to the template
	data := struct {
		Mission models2.Mission
		Heroes  []models2.Hero
	}{
		Mission: mission,
		Heroes:  heroes,
	}

	err = view.Templates.ExecuteTemplate(w, "EditMission.html", data)
	if err != nil {
		log.Println("Erro ao executar template:", err)
		http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
		return
	}
}

// UpdateMission handles the update of a mission's information.
func UpdateMission(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Parse form data
		err := r.ParseForm()
		if err != nil {
			log.Println("Erro ao parsear o formulário:", err)
			http.Error(w, "Erro ao processar o formulário", http.StatusBadRequest)
			return
		}

		// Extract form values
		id, _ := strconv.Atoi(r.FormValue("id"))
		nomeMissao := r.FormValue("nome_missao")
		descricao := r.FormValue("descricao")
		nivelDificuldade, _ := strconv.Atoi(r.FormValue("nivel_dificuldade"))
		recompensa, _ := strconv.Atoi(r.FormValue("recompensa"))
		resultado := r.FormValue("resultado")

		// Get selected heroes (multiple selection)
		heroisDesignadosStr := r.Form["herois_designados"]
		var heroisDesignados []int
		for _, idStr := range heroisDesignadosStr {
			heroID, _ := strconv.Atoi(idStr)
			heroisDesignados = append(heroisDesignados, heroID)
		}

		// Create a Mission object with updated data
		updatedMission := models2.Mission{
			ID:           id,
			MissionName:  nomeMissao,
			Description:  descricao,
			Dificulty:    nivelDificuldade,
			HeroAssigned: heroisDesignados,
			Result:       resultado,
			Reward:       recompensa,
		}

		// Update in database
		db := dataBase.ConnectDataBase()
		defer db.Close()

		err = models2.ModifyMission(db, id, updatedMission)
		if err != nil {
			log.Println("Erro ao atualizar missão:", err)
			http.Error(w, "Erro ao atualizar missão", http.StatusInternalServerError)
			return
		}

		// Redirect to the list of missions
		http.Redirect(w, r, "/missions", http.StatusSeeOther)
	}
}
