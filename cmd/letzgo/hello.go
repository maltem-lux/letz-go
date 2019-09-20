package main

import (
	"github.com/maltem-lux/letz-go/internal/database"
	"github.com/maltem-lux/letz-go/internal/entitites"
	"log"
)



func main() {
	//getAllAbilities()
	log.Println("second call")
	//getAllAbilities()
	//database.DbMgr.ExecuteSelectQuery("")
	c := &entitites.Character{Name:"Bobby"}
	log.Println("Char before : " + c.ToString())
	database.DbMgr.GetConnection().Where("Name = ?", "Bobby").First(&c)
	log.Println("Char after : " + c.ToString())
}