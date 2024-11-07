// controllers/crime_controller.go
package controllers

import (
	"hero-api/dataBase"
	models2 "hero-api/models"
	"hero-api/view"
	"log"
	"net/http"
	"strconv"
)

// ListCrimes displays all crimes.
func ListCrimes(w http.ResponseWriter, r *http.Request) {
	db := dataBase.ConnectDataBase()
	defer db.Close()

	crimes, err := models2.GetAllCrimes(db)
	if err != nil {
		log.Println("Erro ao obter crimes:", err)
		http.Error(w, "Erro ao obter crimes", http.StatusInternalServerError)
		return
	}

	view.Templates.ExecuteTemplate(w, "ListCrimes", crimes)
}

// CreateCrime renders the page to create a new crime.
func CreateCrime(w http.ResponseWriter, r *http.Request) {
	db := dataBase.ConnectDataBase()
	defer db.Close()

	// Get all heroes to select as the responsible hero
	heroes, err := models2.GetAllHeroes(db)
	if err != nil {
		log.Println("Erro ao obter heróis:", err)
		http.Error(w, "Erro ao obter heróis", http.StatusInternalServerError)
		return
	}

	view.Templates.ExecuteTemplate(w, "CreateCrime", heroes)
}

// InsertCrime handles the insertion of a new crime.
func InsertCrime(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Parse form data
		err := r.ParseForm()
		if err != nil {
			log.Println("Erro ao parsear o formulário:", err)
			http.Error(w, "Erro ao processar o formulário", http.StatusBadRequest)
			return
		}

		// Extract form values
		nomeCrime := r.FormValue("nome_crime")
		descricao := r.FormValue("descricao")
		dataCrime := r.FormValue("data_crime")
		heroiResponsavel, _ := strconv.Atoi(r.FormValue("heroi_responsavel"))
		severidade, _ := strconv.Atoi(r.FormValue("severidade"))

		// Create a new Crime object
		newCrime := models2.Crime{
			NomeCrime:        nomeCrime,
			Descricao:        descricao,
			DataCrime:        dataCrime,
			HeroiResponsavel: heroiResponsavel,
			Severidade:       severidade,
		}

		// Insert into database
		db := dataBase.ConnectDataBase()
		defer db.Close()

		err = models2.AddCrime(db, newCrime)
		if err != nil {
			log.Println("Erro ao adicionar crime:", err)
			http.Error(w, "Erro ao adicionar crime", http.StatusInternalServerError)
			return
		}

		// Redirect to the list of crimes
		http.Redirect(w, r, "/crimes", http.StatusSeeOther)
	}
}

// DeleteCrime handles the deletion of a crime.
func DeleteCrime(w http.ResponseWriter, r *http.Request) {
	// Get the crime ID from the URL query
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("ID inválido:", err)
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	db := dataBase.ConnectDataBase()
	defer db.Close()

	err = models2.RemoveCrime(db, id)
	if err != nil {
		log.Println("Erro ao remover crime:", err)
		http.Error(w, "Erro ao remover crime", http.StatusInternalServerError)
		return
	}

	// Redirect to the list of crimes
	http.Redirect(w, r, "/crimes", http.StatusSeeOther)
}

// EditCrime renders the page to edit an existing crime.
func EditCrime(w http.ResponseWriter, r *http.Request) {
	// Get the crime ID from the URL query
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("ID inválido:", err)
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	db := dataBase.ConnectDataBase()
	defer db.Close()

	crime, err := models2.GetCrimeByID(db, id)
	if err != nil {
		log.Println("Erro ao obter crime:", err)
		http.Error(w, "Erro ao obter crime", http.StatusInternalServerError)
		return
	}

	// Get all heroes to select as the responsible hero
	heroes, err := models2.GetAllHeroes(db)
	if err != nil {
		log.Println("Erro ao obter heróis:", err)
		http.Error(w, "Erro ao obter heróis", http.StatusInternalServerError)
		return
	}

	// Pass both crime and heroes to the template
	data := struct {
		Crime  models2.Crime
		Heroes []models2.Hero
	}{
		Crime:  crime,
		Heroes: heroes,
	}

	view.Templates.ExecuteTemplate(w, "EditCrime", data)
}

// UpdateCrime handles the update of a crime's information.
func UpdateCrime(w http.ResponseWriter, r *http.Request) {
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
		nomeCrime := r.FormValue("nome_crime")
		descricao := r.FormValue("descricao")
		dataCrime := r.FormValue("data_crime")
		heroiResponsavel, _ := strconv.Atoi(r.FormValue("heroi_responsavel"))
		severidade, _ := strconv.Atoi(r.FormValue("severidade"))
		oculto := r.FormValue("oculto") == "on"

		// Create a Crime object with updated data
		updatedCrime := models2.Crime{
			ID:               id,
			NomeCrime:        nomeCrime,
			Descricao:        descricao,
			DataCrime:        dataCrime,
			HeroiResponsavel: heroiResponsavel,
			Severidade:       severidade,
			Oculto:           oculto,
		}

		// Update in database
		db := dataBase.ConnectDataBase()
		defer db.Close()

		err = models2.ModifyCrime(db, id, updatedCrime)
		if err != nil {
			log.Println("Erro ao atualizar crime:", err)
			http.Error(w, "Erro ao atualizar crime", http.StatusInternalServerError)
			return
		}

		// Redirect to the list of crimes
		http.Redirect(w, r, "/crimes", http.StatusSeeOther)
	}
}
