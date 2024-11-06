package controllers

import (
	"hero-api/src/models"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

// Renderizar templates
var templates = template.Must(template.ParseGlob("../templates/*.html"))

// Pagina principal
func Index(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "Index", nil)
}

// Pagina para listar herois
func ListHeroes(w http.ResponseWriter, r *http.Request) {
	heroes := GetHeroes()
	templates.ExecuteTemplate(w, "ListHeroes", heroes)
}

// Pagina para criar um heroi
func CreateHero(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "CreateHero", nil)
}

// Pagina para inserir novo heroi
func InsertHero(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		realName := r.FormValue("realName")
		heroName := r.FormValue("heroName")
		gender := r.FormValue("gender")
		height, _ := strconv.ParseFloat(r.FormValue("height"), 64)
		weight, _ := strconv.ParseFloat(r.FormValue("weight"), 64)
		birthDate, _ := time.Parse("2006-01-02", r.FormValue("birthDate"))
		birthPlace := r.FormValue("birthPlace")
		strengthLevel, _ := strconv.Atoi(r.FormValue("strengthLevel"))
		popularity, _ := strconv.Atoi(r.FormValue("popularity"))
		status := r.FormValue("status")

		hero := models.Hero{
			RealName:      realName,
			HeroName:      heroName,
			Gender:        gender,
			Height:        height,
			Weight:        weight,
			BirthDate:     birthDate,
			BirthPlace:    birthPlace,
			StrengthLevel: strengthLevel,
			Popularity:    popularity,
			Status:        status,
		}

		AddHero(hero)
	}
	http.Redirect(w, r, "/", 301)
}

// Pagina para atualizar um heroi
func EditH(w http.ResponseWriter, r *http.Request) {
	idHero := r.URL.Query().Get("id")
	hero := EditHero(idHero)
	templates.ExecuteTemplate(w, "UpdateHero", hero)
}

// Pagina para atualizar um heroi
func UpdateH(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		realName := r.FormValue("realName")
		heroName := r.FormValue("heroName")
		gender := r.FormValue("gender")
		height, _ := strconv.ParseFloat(r.FormValue("height"), 64)
		weight, _ := strconv.ParseFloat(r.FormValue("weight"), 64)
		birthDate, _ := time.Parse("2006-01-02", r.FormValue("birthDate"))
		birthPlace := r.FormValue("birthPlace")
		strengthLevel, _ := strconv.Atoi(r.FormValue("strengthLevel"))
		popularity, _ := strconv.Atoi(r.FormValue("popularity"))
		status := r.FormValue("status")

		hero := models.Hero{
			RealName:      realName,
			HeroName:      heroName,
			Gender:        gender,
			Height:        height,
			Weight:        weight,
			BirthDate:     birthDate,
			BirthPlace:    birthPlace,
			StrengthLevel: strengthLevel,
			Popularity:    popularity,
			Status:        status,
		}
		UpdateHero(hero)
	}
	http.Redirect(w, r, "/", 301)
}

// Pagina para deletar um heroi
func DeleteHero(w http.ResponseWriter, r *http.Request) {
	idHero := r.URL.Query().Get("id")
	RemoveHero(idHero)
	http.Redirect(w, r, "/", 301)
}
