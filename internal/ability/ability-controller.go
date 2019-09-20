package ability

import (
		"fmt"
	"log"
	"net/http"
	"encoding/json"
)

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
	fmt.Println("Endpoint Hit: returnAbilities")
	enableCors(&w)
	a:= FindAll()
	json.NewEncoder(w).Encode(a)
	fmt.Println("Abilities retrieved.")
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}