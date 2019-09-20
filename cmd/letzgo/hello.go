package main

import (
	"fmt"
	"github.com/maltem-lux/letz-go/internal/database"
	"log"
)

func main() {
	//handleRequests()
	getAllEmployees()
	log.Println("second call")
	getAllEmployees()

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