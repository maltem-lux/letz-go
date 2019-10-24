package skills

import "github.com/maltem-lux/letz-go/internal/database"

func FindAll() *[]Skill {
	a := make([]Skill, 0)
	database.DbMgr.GetConnection().Find(&a)
	return &a
}
