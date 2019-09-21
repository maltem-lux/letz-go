package ability

import (
	"encoding/json"
	"fmt"
	"github.com/maltem-lux/letz-go/internal/cors"
	"log"
	"net/http"
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

func returnAbilities(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	fmt.Println("Endpoint Hit: returnAbilities")

	a:= FindAll()
	json.NewEncoder(w).Encode(a)
	fmt.Println("Abilities retrieved.")
}