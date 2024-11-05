package routes

import (
	"github.com/gin-gonic/gin"
	"hero-api/src/controllers"
	"hero-api/src/dataBase"
)

func HandleRequests() {

	dataBase.ConnectDataBase()
	routes := gin.Default()

	routes.GET("/heroes", controllers.GetHeroes)
	routes.GET("/heroes/:id", controllers.GetHero)
	routes.POST("/heroes", controllers.AddHero)
	routes.PUT("/heroes/:id", controllers.UpdateHero)
	routes.DELETE("/heroes/:id", controllers.DeleteHero)
	routes.GET("/heroes/name/:hero_name", controllers.GetHeroByHeroName)

	routes.Run(":8080")
}
