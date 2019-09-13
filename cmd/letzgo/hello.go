package main

import (
	"encoding/json"
	"fmt"
	"github.com/maltem-lux/letz-go/internal/data"
	"log"
	"net/http"
)

var article data.Article = data.Article{Title: "Title of the Article", Desc: "Article Description", Content: "Article Content"}

func main() {
	fmt.Println(article)
	handleRequests()
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