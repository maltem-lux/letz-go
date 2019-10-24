package character

import "github.com/maltem-lux/letz-go/internal/database"

func FindById(id int32) *Character {
	c := &Character{}
	database.DbMgr.GetConnection().Where(CharId + " = ?", id).First(&c)
	return c
}

func FindByPlayerId(id int32) *[]Character {
	c := make([]Character, 0)
	database.DbMgr.GetConnection().Where(PlayerId + " = ?", id).Find(&c)
	return &c
}
