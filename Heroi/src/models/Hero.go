package models

import "time"

type Hero struct {
	ID            int       `json:"id"`
	RealName      string    `json:"real_name"`      // Nome real do herói
	HeroName      string    `json:"hero_name"`      // Nome de herói
	Gender        string    `json:"gender"`         // Sexo
	Height        float64   `json:"height"`         // Altura em metros
	Weight        float64   `json:"weight"`         // Peso em quilogramas
	BirthDate     time.Time `json:"birth_date"`     // Data de nascimento
	BirthPlace    string    `json:"birth_place"`    // Local de nascimento
	StrengthLevel int       `json:"strength_level"` // Nível de força
	Popularity    int       `json:"popularity"`     // Popularidade
	Status        string    `json:"status"`         // Status (e.g., Ativo, Banido)
}
