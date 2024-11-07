package models

import "time"

type Crimes struct {
	ID          int       `json:"id"`
	Hero_id     int       `json:"hero_id"`
	CrimeName   string    `json:"name"`
	Description string    `json:"description"`
	CrimeDate   time.Time `json:"crime_date"`
	severity    int       `json:"severity"`
}
