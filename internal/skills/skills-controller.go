package skills

import (
	"encoding/json"
	"fmt"
	"github.com/maltem-lux/letz-go/internal/cors"
	"net/http"
)

func ReturnSkills(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	fmt.Println("Endpoint Hit: returnSkills")

	a:= FindAll()
	json.NewEncoder(w).Encode(a)
	fmt.Println("Skills retrieved.")
}
