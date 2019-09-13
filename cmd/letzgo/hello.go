package main

import (
	"database/sql"
		"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"encoding/json"
	"github.com/maltem-lux/letz-go/internal/data"
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

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "postgres"
)
var ability Ability
var abilities = make([]Ability, 6)
var article data.Article

func main() {
	handleRequests()
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func connect() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	//return fmt.Sprintf("SELECT * FROM employee WHERE email='%s';", email)

	rows, err := db.Query("SELECT * FROM abilities")
	if err != nil {
		panic(err)
	}
	c := 0
	for rows.Next() {
		rows.Scan(&ability.Ability_id, &ability.Name, &ability.Short_nm, &ability.Description, &ability.Value, &ability.Modifier, &ability.Bonus, &ability.Penalty)
		abilities = append(abilities)
		c++
	}

	fmt.Println("Abilities retrieved.")
}



func handleRequests() {
	http.HandleFunc("/home", homePage)
	http.HandleFunc("/abilities", returnAbilities)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func returnAbilities(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: returnAbilities")
	connect()
	fmt.Println(abilities)
	json.NewEncoder(w).Encode(abilities)
	fmt.Println("Abilities returned.")
}