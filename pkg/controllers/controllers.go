package controllers

import (
	"fmt"
	"github.com/maltem-lux/letz-go/internal/ability"
	"github.com/maltem-lux/letz-go/internal/charAbilities"
	"github.com/maltem-lux/letz-go/internal/charSkills"
	"github.com/maltem-lux/letz-go/internal/character"
	"github.com/maltem-lux/letz-go/internal/skill"
	"log"
	"net/http"
)

func HandleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/characters", character.Handler)
	http.HandleFunc("/newCharacter", character.Create)
	http.HandleFunc("/abilities", ability.ReturnAbilities)
	http.HandleFunc("/charAbilities", charAbilities.ReturnCharAbilities)
	http.HandleFunc("/skills", skill.ReturnSkills)
	http.HandleFunc("/charSkills", charSkills.ReturnCharSkills)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
