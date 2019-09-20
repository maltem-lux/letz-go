package character

import (
	"github.com/maltem-lux/letz-go/internal/database"
)

func FindById(id int32) *Character {
	c := &Character{}
	database.DbMgr.GetConnection().Where(CHAR_ID + " = ?", id).First(&c)
	return c
}

func FindByPlayerId(id int32) *Character {
	c := &Character{}
	database.DbMgr.GetConnection().Where(PLAYER_ID + " = ?", id).First(&c)
	return c
}