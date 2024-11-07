package models

import (
	"database/sql"
	"log"
	"strings"
)

type Hero struct {
	ID            int            `json:"id"`
	RealName      string         `json:"real_name"`      // Nome real do herói
	HeroName      string         `json:"hero_name"`      // Nome de herói
	Gender        string         `json:"gender"`         // Sexo
	Height        float64        `json:"height"`         // Altura em metros
	Weight        float64        `json:"weight"`         // Peso em quilogramas
	BirthDate     string         `json:"birth_date"`     // Data de nascimento
	BirthPlace    string         `json:"birth_place"`    // Local de nascimento
	Powers        []string       `json:"powers"`         // Poderes
	StrengthLevel int            `json:"strength_level"` // Nível de força
	Popularity    int            `json:"popularity"`     // Popularidade
	Status        string         `json:"status"`         // Status (e.g., Ativo, Banido)
	BattleHistory []BattleRecord `json:"battle_history"` // Histórico de batalhas
}

// Funcao para adicionar um novo herois
func AddHero(db *sql.DB, hero Hero) error {

	powersString := strings.Join(hero.Powers, ",")

	insertData, err := db.Prepare("insert into heroes (nome_real, nome_heroi, sexo, altura, peso, data_nascimento, local_nascimento, poderes, nivel_forca, popularidade, status) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)")

	if err != nil {
		log.Println("Erro ao inserir herói (Prepare): ", err)
		return err
	}

	_, err = insertData.Exec(hero.RealName, hero.HeroName, hero.Gender, hero.Height,
		hero.Weight, hero.BirthDate, hero.BirthPlace, powersString, hero.StrengthLevel, hero.Popularity, hero.Status)

	if err != nil {
		log.Println("Erro ao inserir herói (Exec): ", err)
		return err
	}

	return nil
}

// Funcao para deletar um heroi
func RemoveHero(db *sql.DB, id int) error {

	query := "delete from herois where id = $1"

	_, err := db.Exec(query, id)

	if err != nil {
		log.Println("Erro ao deletar herói: ", err)
		return err
	}

	return nil
}

// Funcao para atualizar um heroi
func ModifyHero(db *sql.DB, id int, hero Hero) error {

	powersString := strings.Join(hero.Powers, ",")

	update, err := db.Prepare("update herois set nome_real = $1, nome_heroi = $2, sexo = $3, altura = $4, peso = $5, data_nascimento = $6, local_nascimento = $7, poderes = $8, nivel_forca = $9, popularidade = $10, status = $11 where id = $12")

	if err != nil {
		log.Println("Erro ao atualizar herói (Prepare): ", err)
		return err
	}

	_, err = update.Exec(hero.RealName, hero.HeroName, hero.Gender, hero.Height, hero.Weight, hero.BirthDate, hero.BirthPlace, powersString, hero.StrengthLevel, hero.Popularity, hero.Status, hero.ID)

	if err != nil {
		log.Println("Erro ao atualizar herói (Exec): ", err)
		return err
	}

	return nil
}

// Funcao para obter um heroi pelo ID
func GetHero(sb *sql.DB, parameters string) ([]Hero, error) {
	var heroes []Hero

	query := "select * from herois where nome_heroi ilike '%' || $1 || '%' or status ilike '%' || $1 || '%'"

	rows, err := sb.Query(query, parameters)

	if err != nil {
		log.Println("Erro ao buscar herói: ", err)
		return heroes, err
	}

	defer rows.Close()

	for rows.Next() {
		var hero Hero
		var powersString string

		err := rows.Scan(&hero.ID, &hero.RealName, &hero.HeroName, &hero.Gender, &hero.Height, &hero.Weight,
			&hero.BirthDate, &hero.BirthPlace, &powersString, &hero.StrengthLevel, &hero.Popularity, &hero.Status)

		if err != nil {
			log.Println("Erro ao buscar herói: ", err)
			continue
		}

		hero.Powers = strings.Split(powersString, ",")
		heroes = append(heroes, hero)
	}

	return heroes, nil
}

// Funcao para atualizar a popularidade de um heroi
func UpdateStatusPopularity(db *sql.DB, id int) error {
	hero, err := GetHeroByID(db, id)

	if err != nil {
		return err
	}

	if hero.Popularity < 20 {
		hero.Status = "Banido"
		err = ModifyHero(db, id, hero)
		if err != nil {
			return err
		}
	}
	return nil
}

// Funcao para obter um heroi pelo ID
func GetHeroByID(db *sql.DB, id int) (Hero, error) {
	var hero Hero
	var powersString string

	query := "select * from herois where id = $1"

	row := db.QueryRow(query, id)

	err := row.Scan(&hero.ID, &hero.RealName, &hero.HeroName, &hero.Gender, &hero.Height, &hero.Weight,
		&hero.BirthDate, &hero.BirthPlace, &powersString, &hero.StrengthLevel, &hero.Popularity, &hero.Status)

	if err != nil {
		log.Println("Erro ao buscar herói: ", err)
		return hero, err
	}

	hero.Powers = strings.Split(powersString, ",")

	return hero, nil
}

func GetAllHeroes(db *sql.DB) ([]Hero, error) {
	var heroes []Hero
	query := `
		SELECT id, nome_real, nome_heroi, sexo, altura, peso, data_nascimento, 
		       local_nascimento, poderes, nivel_forca, popularidade, status 
		FROM heroes`
	rows, err := db.Query(query)
	if err != nil {
		log.Println("Erro ao consultar heróis:", err)
		return heroes, err
	}
	defer rows.Close()

	for rows.Next() {
		var hero Hero
		var poderesStr string
		err := rows.Scan(&hero.ID, &hero.RealName, &hero.HeroName, &hero.Gender, &hero.Height, &hero.Weight,
			&hero.BirthDate, &hero.BirthPlace, &poderesStr, &hero.StrengthLevel, &hero.Popularity, &hero.Status)
		if err != nil {
			log.Println("Erro ao ler herói:", err)
			continue
		}
		hero.Powers = strings.Split(poderesStr, ",")
		heroes = append(heroes, hero)
	}
	return heroes, nil
}
