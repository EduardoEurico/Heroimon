// controllers/hero_controller.go
package controllers

import (
	_ "embed"
	"hero-api/dataBase"
	"hero-api/models"
	"hero-api/view"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Index renders the home page.
func Index(w http.ResponseWriter, r *http.Request) {
	view.Templates.ExecuteTemplate(w, "Index", nil)
}

// ListHeroes displays all heroes.
func ListHeroes(w http.ResponseWriter, r *http.Request) {
	db := dataBase.ConnectDataBase()
	defer db.Close()

	heroes, err := models.GetAllHeroes(db)
	if err != nil {
		log.Println("Erro ao obter heróis:", err)
		http.Error(w, "Erro ao obter heróis", http.StatusInternalServerError)
		return
	}

	view.Templates.ExecuteTemplate(w, "ListHeroes", heroes)
}

// CreateHero renders the page to create a new hero.
func CreateHero(w http.ResponseWriter, r *http.Request) {
	view.Templates.ExecuteTemplate(w, "CreateHero", nil)
}

// InsertHero handles the insertion of a new hero.
func InsertHero(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Parse form data
		err := r.ParseForm()
		if err != nil {
			log.Println("Erro ao parsear o formulário:", err)
			http.Error(w, "Erro ao processar o formulário", http.StatusBadRequest)
			return
		}

		// Extract form values
		nomeReal := r.FormValue("nome_real")
		nomeHeroi := r.FormValue("nome_heroi")
		sexo := r.FormValue("sexo")
		altura, _ := strconv.ParseFloat(r.FormValue("altura"), 64)
		peso, _ := strconv.ParseFloat(r.FormValue("peso"), 64)
		dataNascimento := r.FormValue("data_nascimento")
		localNascimento := r.FormValue("local_nascimento")
		poderes := strings.Split(r.FormValue("poderes"), ",") // Poderes separados por vírgula
		nivelForca, _ := strconv.Atoi(r.FormValue("nivel_forca"))
		popularidade, _ := strconv.Atoi(r.FormValue("popularidade"))
		status := r.FormValue("status")

		// Create a new Hero object
		newHero := models.Hero{
			RealName:      nomeReal,
			HeroName:      nomeHeroi,
			Gender:        sexo,
			Height:        altura,
			Weight:        peso,
			BirthDate:     dataNascimento,
			BirthPlace:    localNascimento,
			Powers:        poderes,
			StrengthLevel: nivelForca,
			Popularity:    popularidade,
			Status:        status,
		}

		// Insert into database
		db := dataBase.ConnectDataBase()
		models.AddHero(db, newHero)

		defer db.Close()

		err = models.AddHero(db, newHero)
		if err != nil {
			log.Println("Erro ao adicionar herói:", err)
			http.Error(w, "Erro ao adicionar herói", http.StatusInternalServerError)
			return
		}

		// Redirect to the list of heroes
		http.Redirect(w, r, "/heroes", http.StatusSeeOther)
	}
}

// DeleteHero handles the deletion of a hero.
func DeleteHero(w http.ResponseWriter, r *http.Request) {
	// Get the hero ID from the URL query
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("ID inválido:", err)
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	db := dataBase.ConnectDataBase()
	defer db.Close()

	err = models.RemoveHero(db, id)
	if err != nil {
		log.Println("Erro ao remover herói:", err)
		http.Error(w, "Erro ao remover herói", http.StatusInternalServerError)
		return
	}

	// Redirect to the list of heroes
	http.Redirect(w, r, "/heroes", http.StatusSeeOther)
}

// EditHero renders the page to edit an existing hero.
func EditHero(w http.ResponseWriter, r *http.Request) {
	// Get the hero ID from the URL query
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("ID inválido:", err)
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	db := dataBase.ConnectDataBase()
	defer db.Close()

	hero, err := models.GetHeroByID(db, id)
	if err != nil {
		log.Println("Erro ao obter herói:", err)
		http.Error(w, "Erro ao obter herói", http.StatusInternalServerError)
		return
	}

	view.Templates.ExecuteTemplate(w, "EditHero", hero)
}

// UpdateHero handles the update of a hero's information.
func UpdateHero(w http.ResponseWriter, r *http.Request) {
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
		nomeReal := r.FormValue("nome_real")
		nomeHeroi := r.FormValue("nome_heroi")
		sexo := r.FormValue("sexo")
		altura, _ := strconv.ParseFloat(r.FormValue("altura"), 64)
		peso, _ := strconv.ParseFloat(r.FormValue("peso"), 64)
		dataNascimento := r.FormValue("data_nascimento")
		localNascimento := r.FormValue("local_nascimento")
		poderes := strings.Split(r.FormValue("poderes"), ",")
		nivelForca, _ := strconv.Atoi(r.FormValue("nivel_forca"))
		popularidade, _ := strconv.Atoi(r.FormValue("popularidade"))
		status := r.FormValue("status")

		// Create a Hero object with updated data
		updatedHero := models.Hero{
			RealName:      nomeReal,
			HeroName:      nomeHeroi,
			Gender:        sexo,
			Height:        altura,
			Weight:        peso,
			BirthDate:     dataNascimento,
			BirthPlace:    localNascimento,
			Powers:        poderes,
			StrengthLevel: nivelForca,
			Popularity:    popularidade,
			Status:        status,
		}

		// Update in database
		db := dataBase.ConnectDataBase()
		defer db.Close()

		err = models.ModifyHero(db, id, updatedHero)
		if err != nil {
			log.Println("Erro ao atualizar herói:", err)
			http.Error(w, "Erro ao atualizar herói", http.StatusInternalServerError)
			return
		}

		// Redirect to the list of heroes
		http.Redirect(w, r, "/heroes", http.StatusSeeOther)
	}
}
