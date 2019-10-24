package character

import (
	"encoding/json"
	"fmt"
	"github.com/maltem-lux/letz-go/internal/cors"
	"log"
	"net/http"
)

// TODO : Create here the method to handle all characters with parameters : (w http.ResponseWriter, r *http.Request)
// TODO : Error Handling
// TODO : Tips : To enable Cors, you can call cors.EnableCors(&w), I let you be curious and check what it does.
func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	cors.EnableCors(&w)

	fmt.Println("Endpoint Hit: return AllCharactersOfPlayer")
	a:= FindByPlayerId(1)
	json.NewEncoder(w).Encode(a)
	fmt.Println("Characters retrieved.")
}
