package charAbilities

import (
	"github.com/maltem-lux/letz-go/internal/database"
				)

func FindAll() *[]CharAbilities {
	a := make([]CharAbilities, 0)
	database.DbMgr.GetConnection().Find(&a)
	return &a
}