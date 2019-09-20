package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

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
	json.NewEncoder(w).Encode("")
}