package character

import (
	"encoding/json"
	"fmt"
	"github.com/maltem-lux/letz-go/internal/cors"
	"log"
	"net/http"
)

// TODO : Adapt this method to retrieve the param charId if it exists.
// TODO : Depending on the value of the param
//  either call the findByPlayerId to get all chars
//  or the findById to get his/her details
func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	cors.EnableCors(&w)

	fmt.Println("Endpoint Hit: return AllCharactersOfPlayer")
	a:= FindByPlayerId(1)
	json.NewEncoder(w).Encode(a)
	fmt.Println("Characters retrieved.")
}
