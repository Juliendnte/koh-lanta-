package controller

import (
	"fmt"
	InitStruct "koh-lanta/back"
	InitTemps "koh-lanta/temps"
	"net/http"
	"strconv"
)

var err error

func Index(w http.ResponseWriter, r *http.Request) {
	InitTemps.Temp.ExecuteTemplate(w, "index", nil)
}

func Add(w http.ResponseWriter, r *http.Request) {
	InitStruct.Person.Genders = r.URL.Query().Get("gender")
	if InitStruct.Person.Genders == "h" {
		InitTemps.Temp.ExecuteTemplate(w, "addh", InitStruct.Person)
	} else if InitStruct.Person.Genders == "f" {
		InitTemps.Temp.ExecuteTemplate(w, "addf", InitStruct.Person)
	} else {
		fmt.Println("Erreur gender")
	}
}

func InitAdd(w http.ResponseWriter, r *http.Request) {
	InitStruct.Persons, err = InitStruct.ReadJSON()
	if err != nil {
		fmt.Println("Erreur read", err.Error())
	}
	InitStruct.Person.Id = InitStruct.GenerateID()
	InitStruct.Person.Name = r.FormValue("name")
	InitStruct.Person.Clothes, err = strconv.Atoi(r.URL.Query().Get("image"))
	if err != nil {
		fmt.Println("Erreur atoi", err.Error())
	}

	InitStruct.Persons = append(InitStruct.Persons, InitStruct.Person)
	InitStruct.EditJSON(InitStruct.Persons)

	http.Redirect(w, r, "/index", http.StatusMovedPermanently)
}

func Gender(w http.ResponseWriter, r *http.Request) {
	InitTemps.Temp.ExecuteTemplate(w, "gender", nil)

}
