package main

import (
	"github.com/maltem-lux/letz-go/internal/character"
	"log"
)



func main() {
	//getAllAbilities()
	log.Println("second call")
	//getAllAbilities()
	//database.DbMgr.ExecuteSelectQuery("")
	c1 := character.FindById(1)
	log.Println("Char1 : " + c1.ToString())
}