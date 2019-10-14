package character

import (
	"encoding/json"
	"fmt"
	"github.com/maltem-lux/letz-go/internal/cors"
	"log"
	"net/http"
	"strconv"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	cors.EnableCors(&w)

	keys, ok := r.URL.Query()["charId"]
	if !ok || len(keys[0]) < 1 {
		handleAllCharsOfPlayer(w)
	} else {
		handleCharDetails(w, keys[0])
	}
}

func handleCharDetails(w http.ResponseWriter, paramCharId string) {
	charId,err := strconv.Atoi(paramCharId)
	if err != nil {
		errMsg, _ := fmt.Printf("Error converting param to int : %s", err)
		panic(errMsg)
	} else {
		fmt.Printf("Endpoint Hit: ReturnCharacterById with ID %d", charId)

		c:= FindById(int32(charId))
		json.NewEncoder(w).Encode(c)
		fmt.Printf("Character with ID %d retrieved.", charId)
	}
}

func handleAllCharsOfPlayer(w http.ResponseWriter) {
	log.Println("Url Param 'charId' is missing")
	fmt.Println("Endpoint Hit: returnCharacters")

	a:= FindByPlayerId(1)
	json.NewEncoder(w).Encode(a)
	fmt.Println("Characters retrieved.")
}
