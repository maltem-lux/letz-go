package controllers

import (
	"fmt"
	"github.com/maltem-lux/letz-go/internal/ability"
	"log"
	"net/http"
)

func HandleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/abilities", ability.ReturnAbilities)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
