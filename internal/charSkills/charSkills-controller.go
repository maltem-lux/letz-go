package charSkills

import (
	"encoding/json"
	"fmt"
	"github.com/maltem-lux/letz-go/internal/cors"
	"net/http"
)

func ReturnCharSkills(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	fmt.Println("Endpoint Hit: returnCharSkills")

	a:= FindAll()
	json.NewEncoder(w).Encode(a)
	fmt.Println("CharSkills retrieved.")
}