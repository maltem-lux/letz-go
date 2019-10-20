package charSkills

import (
	"fmt"
)

type CharSkills struct {
	Char_skill_id int32 `json:"charSkillId"`
	Char_id int32 `json:"charId"`
	Skill_id int32 `json:"skillId"`
	Rank int32 `json:"value"`
}

func (c *CharSkills) ToString() string {
	return fmt.Sprintf("Char_skill_id : %v, Char_id : %v, Skill_id : %v, Rank : %v",
		c.Char_skill_id, c.Char_id, c.Skill_id, c.Rank)
}