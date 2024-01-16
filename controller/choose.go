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
	InitTemps.Temp.ExecuteTemplate(w, "add", InitStruct.Person)
}

func InitAdd(w http.ResponseWriter, r *http.Request) {
	InitStruct.Persons, err = InitStruct.ReadJSON()
	if err != nil {
		fmt.Println("Erreur atoi", err.Error())
	}
	InitStruct.Person.Genders = r.FormValue("sexe")
	InitStruct.Person.Id = InitStruct.GenerateID()
	InitStruct.Person.Name = r.FormValue("name")
	InitStruct.Person.Pants, err = strconv.Atoi(r.FormValue("pants"))
	if err != nil {
		fmt.Println("Erreur atoi", err.Error())
	}
	InitStruct.Person.Shirt, err = strconv.Atoi(r.FormValue("shirt"))
	if err != nil {
		fmt.Println("Erreur atoi", err.Error())
	}

	InitStruct.Persons = append(InitStruct.Persons, InitStruct.Person)
	InitStruct.EditJSON(InitStruct.Persons)

	http.Redirect(w, r, "/index", http.StatusMovedPermanently)
}
