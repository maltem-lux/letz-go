package controllers

import (
	"fmt"
	"github.com/maltem-lux/letz-go/internal/abilities"
	"github.com/maltem-lux/letz-go/internal/charAbilities"
	"github.com/maltem-lux/letz-go/internal/charSkills"
	"github.com/maltem-lux/letz-go/internal/character"
	"github.com/maltem-lux/letz-go/internal/skills"
	"log"
	"net/http"
)

func HandleRequests() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/characters", character.Handler)
	http.HandleFunc("/abilities", abilities.ReturnAbilities)
	http.HandleFunc("/charAbilities", charAbilities.ReturnCharAbilities)
	http.HandleFunc("/skills", skills.ReturnSkills)
	http.HandleFunc("/charSkills", charSkills.ReturnCharSkills)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func handleHome(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
