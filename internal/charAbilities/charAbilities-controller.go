package charAbilities

import (
	"encoding/json"
	"fmt"
	"github.com/maltem-lux/letz-go/internal/cors"
	"net/http"
)

func ReturnCharAbilities(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	fmt.Println("Endpoint Hit: returnCharAbilities")

	a:= FindAll()
	json.NewEncoder(w).Encode(a)
	fmt.Println("CharAbilities retrieved.")
}