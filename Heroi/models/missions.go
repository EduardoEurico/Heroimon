package models

import (
	"database/sql"
	"github.com/lib/pq"
	"log"
)

type Mission struct {
	ID           int    `json:"id"`
	MissionName  string `json:"mission_name"`
	Description  string `json:"description"`
	Dificulty    int    `json:"dificulty"`
	HeroAssigned []int  `json:"hero_assigned"`
	Result       string `json:"result"`
	Reward       int    `json:"reward"`
}

func AddMission(db *sql.DB, mission Mission) error {

	insertData, err := db.Prepare("insert into missions (name, description, dificulty, result, reward) values ($1, $2, $3, $4, $5)")

	if err != nil {
		log.Println("Erro ao inserir missão (Prepare): ", err)
		return err
	}

	_, err = insertData.Exec(mission.MissionName, mission.Description, mission.Dificulty, mission.Result, mission.Reward)

	if err != nil {
		log.Println("Erro ao inserir missão (Exec): ", err)
		return err
	}

	for _, heroID := range mission.HeroAssigned {
		err = UpdateHeroAttributesFromMission(db, heroID, mission)
		if err != nil {
			log.Println("Erro ao atribuir missão ao herói: ", err)
			return err
		}
	}

	return nil
}

func UpdateHeroAttributesFromMission(db *sql.DB, heroID int, mission Mission) error {

	hero, err := GetHeroByID(db, heroID)

	if err != nil {
		log.Println("Erro ao buscar herói: ", err)
		return err
	}

	if mission.Result == "success" {
		hero.Popularity += 10
		hero.StrengthLevel += 5
	} else {
		hero.Popularity -= 10
		if hero.Popularity < 0 {
			hero.Popularity = 0
		}
	}

	err = ModifyHero(db, heroID, hero)

	if err != nil {
		log.Println("Erro ao atualizar atributos do herói: ", err)
		return err
	}

	return nil
}

func ConsultMissions(db *sql.DB, Dificulty, heroID int) ([]Mission, error) {
	var missions []Mission

	query := "select * from missions where nivel_dificuldade <= $1 and $2 = any (herois_designados)"

	rows, err := db.Query(query, Dificulty, heroID)

	if err != nil {
		log.Println("Erro ao buscar missões: ", err)
		return missions, err
	}

	defer rows.Close()

	for rows.Next() {
		var mission Mission

		err = rows.Scan(&mission.ID, &mission.MissionName, &mission.Description,
			&mission.Dificulty, &mission.HeroAssigned, &mission.Result, &mission.Reward)

		if err != nil {
			log.Println("Erro ao buscar missões: ", err)
			continue
		}

		missions = append(missions, mission)

	}

	return missions, nil
}

func GetAllMissions(db *sql.DB) ([]Mission, error) {
	var missions []Mission
	query := `
		SELECT id, nome_missao, descricao, nivel_dificuldade, herois_designados, resultado, recompensa
		FROM missions`
	rows, err := db.Query(query)
	if err != nil {
		log.Println("Erro ao consultar missões:", err)
		return missions, err
	}
	defer rows.Close()

	for rows.Next() {
		var mission Mission
		var heroisDesignados []int
		err := rows.Scan(&mission.ID, &mission.MissionName, &mission.Description, &mission.Dificulty,
			&heroisDesignados, &mission.Result, &mission.Reward)
		if err != nil {
			log.Println("Erro ao ler missão:", err)
			continue
		}
		mission.HeroAssigned = heroisDesignados
		missions = append(missions, mission)
	}
	return missions, nil
}

// GetMissionByID retrieves a mission by ID from the database.
func GetMissionByID(db *sql.DB, id int) (Mission, error) {
	var mission Mission
	var heroisDesignados []int
	query := `
		SELECT id, nome_missao, descricao, nivel_dificuldade, herois_designados, resultado, recompensa
		FROM missions WHERE id = $1`
	row := db.QueryRow(query, id)
	err := row.Scan(&mission.ID, &mission.MissionName, &mission.Description, &mission.Dificulty,
		&heroisDesignados, &mission.Result, &mission.Reward)
	if err != nil {
		log.Println("Erro ao obter missão por ID:", err)
		return mission, err
	}
	mission.HeroAssigned = heroisDesignados
	return mission, nil
}

// ModifyMission updates mission attributes in the database.
func ModifyMission(db *sql.DB, id int, mission Mission) error {
	query := `
		UPDATE missions SET 
		nome_missao = $1, 
		descricao = $2, 
		nivel_dificuldade = $3, 
		herois_designados = $4, 
		resultado = $5, 
		recompensa = $6
		WHERE id = $7`
	_, err := db.Exec(query, mission.MissionName, mission.Description, mission.Dificulty,
		pq.Array(mission.HeroAssigned), mission.Result, mission.Reward, id)
	if err != nil {
		log.Println("Erro ao atualizar missão:", err)
		return err
	}
	return nil
}

// RemoveMission removes a mission from the database.
func RemoveMission(db *sql.DB, id int) error {
	query := `DELETE FROM missions WHERE id = $1`
	_, err := db.Exec(query, id)
	if err != nil {
		log.Println("Erro ao remover missão:", err)
		return err
	}
	return nil
}
