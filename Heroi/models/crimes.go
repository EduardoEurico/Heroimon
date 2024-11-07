package models

import (
	"database/sql"
	"log"
)

// Crime represents a crime with all its attributes.
type Crime struct {
	ID               int
	NomeCrime        string
	Descricao        string
	DataCrime        string
	HeroiResponsavel int
	Severidade       int
	Oculto           bool
}

// AddCrime adds a new crime to the database.
func AddCrime(db *sql.DB, crime Crime) error {
	query := `
		INSERT INTO crimes 
		(nome_crime, descricao, data_crime, heroi_responsavel, severidade, oculto)
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := db.Exec(query, crime.NomeCrime, crime.Descricao, crime.DataCrime, crime.HeroiResponsavel, crime.Severidade, crime.Oculto)
	if err != nil {
		log.Println("Erro ao inserir crime:", err)
		return err
	}

	// Impact on hero's popularity
	err = DecreaseHeroPopularity(db, crime.HeroiResponsavel, crime.Severidade)
	if err != nil {
		return err
	}

	return nil
}

// GetAllCrimes retrieves all crimes from the database.
func GetAllCrimes(db *sql.DB) ([]Crime, error) {
	var crimes []Crime
	query := `
		SELECT id, nome_crime, descricao, data_crime, heroi_responsavel, severidade, oculto
		FROM crimes`
	rows, err := db.Query(query)
	if err != nil {
		log.Println("Erro ao consultar crimes:", err)
		return crimes, err
	}
	defer rows.Close()

	for rows.Next() {
		var crime Crime
		err := rows.Scan(&crime.ID, &crime.NomeCrime, &crime.Descricao, &crime.DataCrime,
			&crime.HeroiResponsavel, &crime.Severidade, &crime.Oculto)
		if err != nil {
			log.Println("Erro ao ler crime:", err)
			continue
		}
		crimes = append(crimes, crime)
	}
	return crimes, nil
}

// GetCrimeByID retrieves a crime by ID from the database.
func GetCrimeByID(db *sql.DB, id int) (Crime, error) {
	var crime Crime
	query := `
		SELECT id, nome_crime, descricao, data_crime, heroi_responsavel, severidade, oculto
		FROM crimes WHERE id = $1`
	row := db.QueryRow(query, id)
	err := row.Scan(&crime.ID, &crime.NomeCrime, &crime.Descricao, &crime.DataCrime,
		&crime.HeroiResponsavel, &crime.Severidade, &crime.Oculto)
	if err != nil {
		log.Println("Erro ao obter crime por ID:", err)
		return crime, err
	}
	return crime, nil
}

// ModifyCrime updates crime attributes in the database.
func ModifyCrime(db *sql.DB, id int, crime Crime) error {
	query := `
		UPDATE crimes SET 
		nome_crime = $1, 
		descricao = $2, 
		data_crime = $3, 
		heroi_responsavel = $4, 
		severidade = $5, 
		oculto = $6
		WHERE id = $7`
	_, err := db.Exec(query, crime.NomeCrime, crime.Descricao, crime.DataCrime,
		crime.HeroiResponsavel, crime.Severidade, crime.Oculto, id)
	if err != nil {
		log.Println("Erro ao atualizar crime:", err)
		return err
	}
	return nil
}

// RemoveCrime removes a crime from the database.
func RemoveCrime(db *sql.DB, id int) error {
	query := `DELETE FROM crimes WHERE id = $1`
	_, err := db.Exec(query, id)
	if err != nil {
		log.Println("Erro ao remover crime:", err)
		return err
	}
	return nil
}

// DecreaseHeroPopularity decreases the hero's popularity based on crime severity.
func DecreaseHeroPopularity(db *sql.DB, heroID, severidade int) error {
	hero, err := GetHeroByID(db, heroID)
	if err != nil {
		return err
	}
	decremento := severidade * 2 // Example logic
	hero.Popularity -= decremento
	if hero.Popularity < 0 {
		hero.Popularity = 0
	}
	err = ModifyHero(db, heroID, hero)
	if err != nil {
		return err
	}
	return nil
}

// ConsultCrimes returns crimes filtered by hero and severity.
func ConsultCrimes(db *sql.DB, heroiID, severidade int) ([]Crime, error) {
	var crimes []Crime
	query := `
		SELECT id, nome_crime, descricao, data_crime, heroi_responsavel, severidade, oculto
		FROM crimes
		WHERE heroi_responsavel = $1 AND severidade >= $2 AND oculto = FALSE`
	rows, err := db.Query(query, heroiID, severidade)
	if err != nil {
		log.Println("Erro ao consultar crimes:", err)
		return crimes, err
	}
	defer rows.Close()

	for rows.Next() {
		var crime Crime
		err := rows.Scan(&crime.ID, &crime.NomeCrime, &crime.Descricao, &crime.DataCrime,
			&crime.HeroiResponsavel, &crime.Severidade, &crime.Oculto)
		if err != nil {
			log.Println("Erro ao ler crime:", err)
			continue
		}
		crimes = append(crimes, crime)
	}
	return crimes, nil
}

// HideCrime marks old crimes as hidden.
func HideCrime(db *sql.DB, id int) error {
	query := `UPDATE crimes SET oculto = TRUE WHERE id = $1`
	_, err := db.Exec(query, id)
	if err != nil {
		log.Println("Erro ao ocultar crime:", err)
		return err
	}
	return nil
}
