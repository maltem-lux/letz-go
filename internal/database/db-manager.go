package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
)

type DbManagerInterface interface {
	GetConnection() *gorm.DB
	// Add other methods
}

type databaseManager struct {
	db *gorm.DB
}

var DbMgr DbManagerInterface

func init() {
	dbConnectionString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		Host, Port, User, Password, DbName)
	con, err := gorm.Open("postgres", dbConnectionString)

	if err != nil {
		panic("Error initializing db")
	} else {
		log.Print("DB Initialized successfully")
	}
	DbMgr = &databaseManager{db: con}
}

func (mgr *databaseManager) GetConnection() *gorm.DB {
	return mgr.db
}