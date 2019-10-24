package database

import "github.com/jinzhu/gorm"

type DbManagerInterface interface {
	GetConnection() *gorm.DB
}

type databaseManager struct {
	db *gorm.DB
}

var DbMgr DbManagerInterface

// TODO : Create a init() function which must connect to DB using the params into db-params.go file
// TODO : It must instantiate a DbMgr object with the db connection
// TODO : There must be error handling



// TODO : In order to extend the DbManagerInterface Interface,
//  Create a function GetConnection() which
//  - will be applied only to a databaseManager object
//  - will return a gorm.DB pointer
