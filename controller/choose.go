package controller

import (
	"fmt"
	InitStruct "koh-lanta/back"
	InitTemps "koh-lanta/temps"
	"net/http"
	"strconv"
)

var err error

// Execute le template Index
func Index(w http.ResponseWriter, r *http.Request) {
	InitTemps.Temp.ExecuteTemplate(w, "index", nil)
}

// Execute le template Gender
func Gender(w http.ResponseWriter, r *http.Request) {
	InitTemps.Temp.ExecuteTemplate(w, "gender", nil)
}

// Execute le template pour ajouter un homme ou une femme selon ce que l'utilisateur a choisi
func Add(w http.ResponseWriter, r *http.Request) {
	InitStruct.Person.Genders = r.URL.Query().Get("gender")
	if InitStruct.Person.Genders == "h" {
		InitTemps.Temp.ExecuteTemplate(w, "addh", InitStruct.Person)
	} else if InitStruct.Person.Genders == "f" {
		InitTemps.Temp.ExecuteTemplate(w, "addf", InitStruct.Person)
	} else {
		fmt.Println("Erreur gender")
		return
	}
}

// Fonction treatment qui ajoute le personnage de l'utilisateur
func InitAdd(w http.ResponseWriter, r *http.Request) {
	InitStruct.Persons, err = InitStruct.ReadJSON()
	if err != nil {
		fmt.Println("Erreur read", err.Error())
	}

	InitStruct.Person.Id = InitStruct.GenerateID() //Génére un Id aléatoire
	InitStruct.Person.Name = r.FormValue("name")
	InitStruct.Person.Age, err = strconv.Atoi(r.FormValue("age"))
	if err != nil {
		fmt.Println("Erreur atoi", err.Error())
	}
	InitStruct.Person.Clothes, err = strconv.Atoi(r.URL.Query().Get("image"))
	if err != nil {
		fmt.Println("Erreur atoi", err.Error())
	}

	InitStruct.Persons = append(InitStruct.Persons, InitStruct.Person)
	InitStruct.EditJSON(InitStruct.Persons)

	http.Redirect(w, r, "/index", http.StatusMovedPermanently)
}
