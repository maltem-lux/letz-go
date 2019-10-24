package charSkills

import "github.com/maltem-lux/letz-go/internal/database"

func FindAll() *[]CharSkills {
	a := make([]CharSkills, 0)
	database.DbMgr.GetConnection().Find(&a)
	return &a
}
