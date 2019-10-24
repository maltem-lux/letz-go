package skills

import (
	"fmt"
)

type Skill struct {
	Skill_id int32 `json:"skillId"`
	Name string `json:"name"`
	Primary_stat string `json:"primaryStat"`
	Description string `json:"description"`
}

func (s *Skill) ToString() string {
	return fmt.Sprintf("Skill_id : %v, Name : %v, Primary_stat : %v, Description : %v",
		s.Skill_id, s.Name, s.Primary_stat, s.Description)
}
