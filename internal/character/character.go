package character

import (
	"fmt"
	"time"
)

const (
	CharId = "Char_id"
	PlayerId = "Player_id"
)

type Character struct {
	Char_id int32 `json:"charId"`
	Race_id int32 `json:"raceId"`
	Player_id int32 `json:"playerId"`
	Campaign_id int32 `json:"campaign"`
	Campaign_credentials int32 `json:"campaignCredentials"`
	Age int32 `json:"age"`
	Gender string `json:"gender"`
	Name string `json:"name"`
	Level int32 `json:"level"`
	Alignment string `json:"alignment"`
	Deity string `json:"deity"`
	Height int32 `json:"height,string"`
	Height_unit string `json:"heightUnit"`
	Weight int32 `json:"weight,string"`
	Weight_unit string `json:"weightUnit"`
	Creation_date time.Time `json:"creationDate"`
}

func (c *Character) ToString() string {
	return fmt.Sprintf("Char_id : %v, Player_id : %v, Age : %v, Gender : %v, Level : %v",
		c.Char_id, c.Player_id, c.Age, c.Gender, c.Level)
}
