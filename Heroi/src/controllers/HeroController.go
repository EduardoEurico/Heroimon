package controllers

import (
	"database/sql"
	"fmt"
	"hero-api/src/dataBase"
	"hero-api/src/models"
	"log"
	"time"
)

// Funcao para obter todos os herois
func GetHeroes() []models.Hero {

	dataBase := dataBase.ConnectDataBase()

	selectAllHerois, err := dataBase.Query("select * from herois order by id asc")

	if err != nil {
		log.Panic("Erro ao buscar lista de herois: ", err)
	}

	hero := models.Hero{}
	heroes := []models.Hero{}

	for selectAllHerois.Next() {
		var (
			id, popularity, StrengthLevel                  int
			realName, heroName, gender, status, birthPlace string
			height, weight                                 float64
			birthDate                                      time.Time
		)

		err = selectAllHerois.Scan(&id, &realName, &heroName, &gender, &height, &weight, &birthDate, &birthPlace, &StrengthLevel, &popularity, &status)

		if err != nil {
			log.Panic("Erro ao buscar lista de herois: ", err)
		}

		hero.ID = id
		hero.RealName = realName
		hero.HeroName = heroName
		hero.Gender = gender
		hero.Height = height
		hero.Weight = weight
		hero.BirthDate = birthDate
		hero.BirthPlace = birthPlace
		hero.StrengthLevel = StrengthLevel
		hero.Popularity = popularity
		hero.Status = status

		heroes = append(heroes, hero)
	}

	defer dataBase.Close()

	return heroes
}

// Funcao para obter um heroi pelo ID
func GetHero(id int) (models.Hero, error) {
	dataBase := dataBase.ConnectDataBase()
	defer dataBase.Close()

	query := `
        SELECT id, nome_real, nome_heroi, sexo, altura, peso, data_nascimento, 
               local_nascimento, nivel_forca, popularidade, status
        FROM herois
        WHERE id = $1
    `

	var hero models.Hero

	// Variáveis temporárias para armazenar os valores retornados
	var (
		idDB, popularity, strengthLevel                int
		realName, heroName, gender, status, birthPlace string
		height, weight                                 float64
		birthDate                                      time.Time
	)

	// Executa a consulta
	err := dataBase.QueryRow(query, id).Scan(
		&idDB,
		&realName,
		&heroName,
		&gender,
		&height,
		&weight,
		&birthDate,
		&birthPlace,
		&strengthLevel,
		&popularity,
		&status,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return hero, fmt.Errorf("Herói não encontrado")
		}
		return hero, err
	}

	// Preenche o struct hero com os dados obtidos
	hero.ID = idDB
	hero.RealName = realName
	hero.HeroName = heroName
	hero.Gender = gender
	hero.Height = height
	hero.Weight = weight
	hero.BirthDate = birthDate
	hero.BirthPlace = birthPlace
	hero.StrengthLevel = strengthLevel
	hero.Popularity = popularity
	hero.Status = status

	return hero, nil
}

// Funcao para adicionar um novo heroi
func AddHero(hero models.Hero) {
	dataBase := dataBase.ConnectDataBase()

	insertData, err := dataBase.Prepare("insert into herois (nome_real, nome_heroi, sexo, altura, peso, data_nascimento, local_nascimento, nivel_forca, popularidade, status) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)")

	if err != nil {
		log.Panic("Erro ao inserir herói: ", err)
	}

	insertData.Exec(hero.RealName, hero.HeroName, hero.Gender, hero.Height, hero.Weight, hero.BirthDate, hero.BirthPlace, hero.StrengthLevel, hero.Popularity, hero.Status)

	defer dataBase.Close()
}

// Funcao para atualizar um heroi
func EditHero(id string) models.Hero {
	dataBase := dataBase.ConnectDataBase()

	heroDB, err := dataBase.Query("select * from herois where id = $1", id)

	if err != nil {
		log.Panic("Erro ao buscar herói: ", err)
	}

	heroUpdate := models.Hero{}

	for heroDB.Next() {
		var (
			id, popularity, StrengthLevel                  int
			realName, heroName, gender, status, birthPlace string
			height, weight                                 float64
			birthDate                                      time.Time
		)

		err = heroDB.Scan(&id, &realName, &heroName, &gender, &height, &weight, &birthDate, &birthPlace, &StrengthLevel, &popularity, &status)

		if err != nil {
			log.Panic("Erro ao buscar herói: ", err)
		}

		heroUpdate.ID = id
		heroUpdate.RealName = realName
		heroUpdate.HeroName = heroName
		heroUpdate.Gender = gender
		heroUpdate.Height = height
		heroUpdate.Weight = weight
		heroUpdate.BirthDate = birthDate
		heroUpdate.BirthPlace = birthPlace
		heroUpdate.StrengthLevel = StrengthLevel
		heroUpdate.Popularity = popularity
		heroUpdate.Status = status
	}

	defer dataBase.Close()

	return heroUpdate
}

// Funcao para atualizar um heroi
func UpdateHero(hero models.Hero) {
	dataBase := dataBase.ConnectDataBase()

	update, err := dataBase.Prepare("update herois set nome_real = $1, nome_heroi = $2, sexo = $3, altura = $4, peso = $5, data_nascimento = $6, local_nascimento = $7, nivel_forca = $8, popularidade = $9, status = $10 where id = $11")

	if err != nil {
		log.Panic("Erro ao atualizar herói: ", err)
	}

	update.Exec(hero.RealName, hero.HeroName, hero.Gender, hero.Height, hero.Weight, hero.BirthDate, hero.BirthPlace, hero.StrengthLevel, hero.Popularity, hero.Status, hero.ID)

	defer dataBase.Close()
}

// Funcao para deletar um heroi
func RemoveHero(id string) {
	dataBase := dataBase.ConnectDataBase()

	delet, err := dataBase.Prepare("delete from herois where id = $1")

	if err != nil {
		log.Panic("Erro ao deletar herói: ", err)
	}

	delet.Exec(id)

	defer dataBase.Close()
}

// Funcao para buscar pelo nome do heroi
func GetHeroByHeroName(name string) ([]models.Hero, error) {
	dataBase := dataBase.ConnectDataBase()
	defer dataBase.Close()

	query := `
        SELECT id, nome_real, nome_heroi, sexo, altura, peso, data_nascimento, 
               local_nascimento, nivel_forca, popularidade, status
        FROM herois
        WHERE nome_heroi ILIKE $1
        ORDER BY id ASC
    `

	// Executa a consulta com o nome fornecido
	rows, err := dataBase.Query(query, "%"+name+"%")
	if err != nil {
		return nil, fmt.Errorf("Erro ao buscar heróis: %v", err)
	}
	defer rows.Close()

	var heroes []models.Hero

	for rows.Next() {
		var hero models.Hero

		// Variáveis temporárias para armazenar os valores retornados
		var (
			idDB, popularity, strengthLevel                  int
			realName, heroNameDB, gender, status, birthPlace string
			height, weight                                   float64
			birthDate                                        time.Time
		)

		// Lê os dados da linha atual
		err = rows.Scan(
			&idDB,
			&realName,
			&heroNameDB,
			&gender,
			&height,
			&weight,
			&birthDate,
			&birthPlace,
			&strengthLevel,
			&popularity,
			&status,
		)

		if err != nil {
			return nil, fmt.Errorf("Erro ao ler dados de herói: %v", err)
		}

		// Preenche o struct hero com os dados obtidos
		hero.ID = idDB
		hero.RealName = realName
		hero.HeroName = heroNameDB
		hero.Gender = gender
		hero.Height = height
		hero.Weight = weight
		hero.BirthDate = birthDate
		hero.BirthPlace = birthPlace
		hero.StrengthLevel = strengthLevel
		hero.Popularity = popularity
		hero.Status = status

		heroes = append(heroes, hero)
	}

	return heroes, nil
}
