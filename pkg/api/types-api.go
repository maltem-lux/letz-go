package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/maltem-lux/letz-go/internal/database"
	"github.com/maltem-lux/letz-go/internal/data"
)

var abilities = make([]data.Ability, 0)

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func HandleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/abilities", returnAbilities)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func returnAbilities(w http.ResponseWriter, r *http.Request){
	abilities = make([]data.Ability, 0)
	enableCors(&w)
	fmt.Println("Endpoint Hit: returnAbilities")
	database.DbMgr.GetConnection().Find(&abilities)
	fmt.Println("Abilities retrieved.")
	json.NewEncoder(w).Encode(abilities)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}