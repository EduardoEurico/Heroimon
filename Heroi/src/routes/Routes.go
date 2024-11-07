package routes

import (
	"hero-api/src/controllers"
	"log"
	"net/http"
)

func HandleRequests() {

	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/heroes", controllers.ListHeroes)
	//http.HandleFunc("/heroes/:id")
	//http.HandleFunc("/heroes/:names-hero")
	http.HandleFunc("/create-hero", controllers.CreateHero)
	http.HandleFunc("/insert-hero", controllers.InsertHero)
	http.HandleFunc("/delete-hero", controllers.DeleteHero)
	http.HandleFunc("/edit-hero", controllers.EditH)
	http.HandleFunc("/update-hero", controllers.UpdateH)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
