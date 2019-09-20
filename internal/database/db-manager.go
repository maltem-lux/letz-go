package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "postgres"
)

type DbManagerInterface interface {
	ExecuteSelectQuery(query string) *sql.Rows
	// Add other methods
}

type databaseManager struct {
	db *sql.DB
}

var DbMgr DbManagerInterface

func init() {
	dbConnectionString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	con, err := sql.Open("postgres", dbConnectionString)

	if err != nil {
		panic("Error initializing db")
	} else {
		log.Print("DB Initialized successfully")
	}
	DbMgr = &databaseManager{db: con}
}

func (mgr *databaseManager) ExecuteSelectQuery(query string) *sql.Rows {

	rows, err := mgr.db.Query(query)
	if err != nil {
		panic(err)
	}
	return rows
}