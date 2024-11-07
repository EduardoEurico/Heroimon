package models

type Missions struct {
	ID          int    `json:"id"`
	MissionName string `json:"mission_name"`
	Description string `json:"description"`
	Dificulty   int    `json:"dificulty"`
	Result      string `json:"result"`
	Reward      string `json:"reward"`
}
