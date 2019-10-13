package charAbilities

import (
	"fmt"
)

type CharAbilities struct {
	Char_ability_id int32 `json:"charAbilityId"`
	Char_id int32 `json:"charId"`
	Ability_id int32 `json:"abilityId"`
	Value int32 `json:"value"`
	Modifier int32 `json:"modifier"`
	Bonus int32 `json:"bonus"`
	Penalty int32 `json:"penalty"`
}

func (c *CharAbilities) ToString() string {
	return fmt.Sprintf("Char_ability_id : %v, Char_id : %v, Ability_id : %v, Value : %v, Modifier: %v, Bonus: %v, Penalty: %v",
		c.Char_ability_id, c.Char_id, c.Ability_id, c.Value, c.Modifier, c.Bonus, c.Penalty)
}