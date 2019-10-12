package ability

import (
	"encoding/json"
	"fmt"
	"github.com/maltem-lux/letz-go/internal/cors"
	"net/http"
)

func ReturnAbilities(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	fmt.Println("Endpoint Hit: returnAbilities")

	a:= FindAll()
	json.NewEncoder(w).Encode(a)
	fmt.Println("Abilities retrieved.")
}