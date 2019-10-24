package controllers

import (
	"fmt"
	"github.com/maltem-lux/letz-go/internal/character"
	"log"
	"net/http"
)

func HandleRequests() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/characters", character.Handler)
	http.HandleFunc("/newCharacter", character.Create)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func handleHome(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
