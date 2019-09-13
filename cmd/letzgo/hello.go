package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/maltem-lux/letz-go/internal/data"
	"log"
	"net/http"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "jerem"
	password = "root"
	dbname   = "hubmaltem"
)

var article data.Article = data.Article{Title: "Title of the Article", Desc: "Article Description", Content: "Article Content"}

func main() {
	fmt.Println(article)
	//handleRequests()
	connect()
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles/1", returnArticle)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func returnArticle(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: returnArticle")
	json.NewEncoder(w).Encode(article)
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
	fmt.Println("hello")

	rows, err := db.Query("SELECT * FROM Employee")
	if err != nil {
		panic(err)
	}

	var col1 string
	var col2 string
	for rows.Next() {
		rows.Scan(&col1, &col2)
		fmt.Println(col1, col2)
	}

	fmt.Println("Successfully connected!")
}