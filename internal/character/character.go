package character

import (
	"fmt"
	"time"
)

const (
	CharId = "Char_id"
	RaceId = "Race_id"
	PlayerId = "Player_id"
	CampaignId = "Campaign_id"
	CampaignCredentials = "Campaign_credentials"
	Age = "Age"
	Gender = "Gender"
	Name = "Name"
	Level = "Level"
	Alignment = "Alignment"
	Deity = "Deity"
	Height = "Height"
	HeightUnit = "Height_unit"
	Weight = "Weight"
	WeightUnit = "WeightUnit"
	Creation_date = "Creation_date"
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
	Height int32 `json:"height"`
	Height_unit string `json:"heightUnit"`
	Weight int32 `json:"weight"`
	Weight_unit string `json:"weightUnit"`
	Creation_date time.Time `json:"creationDate"`
}

func (c *Character) ToString() string {
	return fmt.Sprintf("Char_id : %v, Player_id : %v, Age : %v, Gender : %v, Level : %v",
		c.Char_id, c.Player_id, c.Age, c.Gender, c.Level)
}