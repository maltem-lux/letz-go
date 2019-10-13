package character

import (
	"encoding/json"
	"fmt"
	"github.com/maltem-lux/letz-go/internal/cors"
	"net/http"
)

func ReturnCharacters(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	fmt.Println("Endpoint Hit: returnCharacters")

	a:= FindByPlayerId(1)
	json.NewEncoder(w).Encode(a)
	fmt.Println("Characters retrieved.")
}