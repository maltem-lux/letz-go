package ability

import (
	"github.com/maltem-lux/letz-go/internal/database"
				)

func FindAll() *[]Ability {
	a := make([]Ability, 0)
	database.DbMgr.GetConnection().Find(&a)
	return &a
}