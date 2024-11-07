package routes

import (
	controllers2 "hero-api/controllers"
	"log"
	"net/http"
	"os"
)

func init() {
	err := os.Chdir("/home/kaynan/Documentos/desenvolvimento/go/Heroimon/Heroi/view")
	if err != nil {
		log.Fatalf("Erro ao definir o diretório de trabalho: %v", err)
	}
}

func HandleRequests() {
	http.HandleFunc("/", controllers2.Index)

	// Hero Routes
	http.HandleFunc("/heroes", controllers2.ListHeroes)
	http.HandleFunc("/create-hero", controllers2.CreateHero)
	http.HandleFunc("/insert-hero", controllers2.InsertHero)
	http.HandleFunc("/delete-hero", controllers2.DeleteHero)
	http.HandleFunc("/edit-hero", controllers2.EditHero)
	http.HandleFunc("/update-hero", controllers2.UpdateHero)

	// Mission Routes
	http.HandleFunc("/missions", controllers2.ListMissions)
	http.HandleFunc("/create-mission", controllers2.CreateMission)
	http.HandleFunc("/insert-mission", controllers2.InsertMission)
	http.HandleFunc("/delete-mission", controllers2.DeleteMission)
	http.HandleFunc("/edit-mission", controllers2.EditMission)
	http.HandleFunc("/update-mission", controllers2.UpdateMission)

	// Crime Routes
	http.HandleFunc("/crimes", controllers2.ListCrimes)
	http.HandleFunc("/create-crime", controllers2.CreateCrime)
	http.HandleFunc("/insert-crime", controllers2.InsertCrime)
	http.HandleFunc("/delete-crime", controllers2.DeleteCrime)
	http.HandleFunc("/edit-crime", controllers2.EditCrime)
	http.HandleFunc("/update-crime", controllers2.UpdateCrime)

	// Battle Routes
	http.HandleFunc("/battles", controllers2.ListBattles)
	http.HandleFunc("/start-battle", controllers2.StartBattle)
	http.HandleFunc("/battle-result", controllers2.BattleResult)
	http.HandleFunc("/battle-details", controllers2.BattleDetails)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
