// models/battle_record.go
package models

import (
	"database/sql"
	"log"
	"math/rand"
	"time"
)

// BattleRecord represents a battle between two heroes.
type BattleRecord struct {
	ID         int
	Heroi1ID   int
	Heroi1Nome string
	Heroi2ID   int
	Heroi2Nome string
	Data       time.Time
	Resultado  string
}

// StartBattle simulates a battle between two heroes.
func StartBattle(db *sql.DB, heroi1ID, heroi2ID int) (string, error) {
	heroi1, err := GetHeroByID(db, heroi1ID)
	if err != nil {
		return "", err
	}
	heroi2, err := GetHeroByID(db, heroi2ID)
	if err != nil {
		return "", err
	}

	resultado := CalculateBattleResult(heroi1, heroi2)

	// Update battle records
	now := time.Now()
	battleRecord := BattleRecord{
		Heroi1ID:  heroi1ID,
		Heroi2ID:  heroi2ID,
		Data:      now,
		Resultado: resultado,
	}

	err = AddBattleRecord(db, battleRecord)
	if err != nil {
		return "", err
	}

	return resultado, nil
}

// CalculateBattleResult determines the winner based on strength, popularity, and randomness.
func CalculateBattleResult(heroi1, heroi2 Hero) string {
	rand.Seed(time.Now().UnixNano())
	fatorAleatorio := rand.Intn(20) - 10 // Random factor between -10 and 10

	score1 := heroi1.StrengthLevel + heroi1.Popularity/2 + fatorAleatorio
	score2 := heroi2.StrengthLevel + heroi2.Popularity/2 - fatorAleatorio

	if score1 > score2 {
		return heroi1.HeroName + " venceu!"
	} else if score2 > score1 {
		return heroi2.HeroName + " venceu!"
	} else {
		return "Empate!"
	}
}

// AddBattleRecord adds a battle record to the database.
func AddBattleRecord(db *sql.DB, record BattleRecord) error {
	query := `
		INSERT INTO battle_records 
		(heroi1_id, heroi2_id, data, resultado)
		VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(query, record.Heroi1ID, record.Heroi2ID, record.Data, record.Resultado)
	if err != nil {
		log.Println("Erro ao inserir registro de batalha:", err)
		return err
	}
	return nil
}

// GetAllBattleRecords retrieves all battle records from the database.
func GetAllBattleRecords(db *sql.DB) ([]BattleRecord, error) {
	var battles []BattleRecord
	query := `
		SELECT br.id, br.heroi1_id, h1.nome_heroi, br.heroi2_id, h2.nome_heroi, br.data, br.resultado
		FROM battle_records br
		JOIN heroes h1 ON br.heroi1_id = h1.id
		JOIN heroes h2 ON br.heroi2_id = h2.id
		ORDER BY br.data DESC`
	rows, err := db.Query(query)
	if err != nil {
		log.Println("Erro ao consultar registros de batalha:", err)
		return battles, err
	}
	defer rows.Close()

	for rows.Next() {
		var battle BattleRecord
		err := rows.Scan(&battle.ID, &battle.Heroi1ID, &battle.Heroi1Nome, &battle.Heroi2ID, &battle.Heroi2Nome,
			&battle.Data, &battle.Resultado)
		if err != nil {
			log.Println("Erro ao ler registro de batalha:", err)
			continue
		}
		battles = append(battles, battle)
	}
	return battles, nil
}

// GetBattleRecordByID retrieves a battle record by ID from the database.
func GetBattleRecordByID(db *sql.DB, id int) (BattleRecord, error) {
	var battle BattleRecord
	query := `
		SELECT br.id, br.heroi1_id, h1.nome_heroi, br.heroi2_id, h2.nome_heroi, br.data, br.resultado
		FROM battle_records br
		JOIN heroes h1 ON br.heroi1_id = h1.id
		JOIN heroes h2 ON br.heroi2_id = h2.id
		WHERE br.id = $1`
	row := db.QueryRow(query, id)
	err := row.Scan(&battle.ID, &battle.Heroi1ID, &battle.Heroi1Nome, &battle.Heroi2ID, &battle.Heroi2Nome,
		&battle.Data, &battle.Resultado)
	if err != nil {
		log.Println("Erro ao obter registro de batalha por ID:", err)
		return battle, err
	}
	return battle, nil
}
