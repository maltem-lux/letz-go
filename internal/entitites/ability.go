package entitites

import (
	"fmt"
)

type Ability struct {
	Ability_id int32 `json:"charId"`
	Name string `json:"charId"`
	Short_nm string `json:"charId"`
	Description string `json:"charId"`
}

func (a *Ability) ToString() string {
	return fmt.Sprintf("Ability_id : %v, Name : %v, Short_nm : %v, Description : %v, ",
		a.Ability_id, a.Name, a.Short_nm, a.Description)
}