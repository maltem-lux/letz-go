package ability

import (
	"fmt"
)

type Ability struct {
	Ability_id int32 `json:"abilityId"`
	Name string `json:"name"`
	Short_nm string `json:"shortNm"`
	Description string `json:"description"`
}

func (a *Ability) ToString() string {
	return fmt.Sprintf("Ability_id : %v, Name : %v, Short_nm : %v, Description : %v",
		a.Ability_id, a.Name, a.Short_nm, a.Description)
}