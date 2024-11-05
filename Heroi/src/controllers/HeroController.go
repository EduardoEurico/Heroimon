package controllers

import (
	"github.com/gin-gonic/gin"
	"hero-api/src/dataBase"
	"hero-api/src/models"
	"net/http"
)

// Funcao para obter todos os herois
func GetHeroes(c *gin.Context) {
	var heroes []models.Hero
	dataBase.DB.Find(&heroes)
	c.JSON(http.StatusOK, heroes)
}

// Funcao para obter um heroi pelo ID
func GetHero(c *gin.Context) {
	var hero models.Hero
	id := c.Param("id")

	if err := dataBase.DB.First(&hero, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Heroi nao encontrado"})
		return
	}

	c.JSON(http.StatusOK, hero)
}

// Funcao para adicionar um novo heroi
func AddHero(c *gin.Context) {
	var hero models.Hero
	if err := c.ShouldBindJSON(&hero); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dataBase.DB.Create(&hero)
	c.JSON(http.StatusCreated, hero)
}

// Funcao para atualizar um heroi
func UpdateHero(c *gin.Context) {
	var hero models.Hero
	id := c.Param("id")

	if err := dataBase.DB.First(&hero, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Heroi nao encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&hero); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dataBase.DB.Save(&hero)
	c.JSON(http.StatusOK, hero)
}

// Funcao para deletar um heroi
func DeleteHero(c *gin.Context) {
	var hero models.Hero
	id := c.Param("id")

	if err := dataBase.DB.First(&hero, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Heroi nao encontrado"})
		return
	}

	dataBase.DB.Delete(&hero)
	c.JSON(http.StatusNoContent, nil)
}

// Funcao para buscar pelo nome do heroi
func GetHeroByHeroName(c *gin.Context) {
	var heroes []models.Hero
	name := c.Query("name")

	dataBase.DB.Where("hero_name LIKE ?", "%"+name+"%").Find(&heroes)
	c.JSON(http.StatusOK, heroes)
}
