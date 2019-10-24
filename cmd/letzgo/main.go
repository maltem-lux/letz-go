package main

import "github.com/maltem-lux/letz-go/internal/database"

func main() {

	// TODO Connect to the Database using the init function into db-manager.go file
	database.DbMgr.GetConnection()
}
