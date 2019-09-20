package main

import (
	"fmt"
	"github.com/maltem-lux/letz-go/internal/data"
	"github.com/maltem-lux/letz-go/internal/database"
	"log"
)

type Ability struct {
	Ability_id int
	Name string
	Short_nm string
	Description string
	Value int
	Modifier int
	Bonus int
	Penalty int
}

var ability Ability
var abilities = make([]Ability, 6)
var article data.Article

func main() {
	getAllEmployees()
	log.Println("second call")
	getAllEmployees()
}

func getAllAbilities() {
	rows := database.DbMgr.ExecuteSelectQuery(`SELECT * FROM abilities`)
	for rows.Next() {
		rows.Scan(&ability.Ability_id, &ability.Name, &ability.Short_nm, &ability.Description, &ability.Value, &ability.Modifier, &ability.Bonus, &ability.Penalty)
		abilities = append(abilities)
	}

	fmt.Println("Abilities retrieved.")

}

func getAllEmployees() {
	rows := database.DbMgr.ExecuteSelectQuery(`select "email", "firstName" from "Employee"`)

	var col1 string
	var col2 string
	for rows.Next() {
		rows.Scan(&col1, &col2)
		fmt.Println(col1, col2)
	}
}