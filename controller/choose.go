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
	} else {
		InitTemps.Temp.ExecuteTemplate(w, "addf", InitStruct.Person)
	}
}

func InitAdd(w http.ResponseWriter, r *http.Request) {
	InitStruct.Persons, err = InitStruct.ReadJSON()
	if err != nil {
		fmt.Println("Erreur read", err.Error())
	}
	InitStruct.Person.Id = InitStruct.GenerateID()
	InitStruct.Person.Name = r.FormValue("name")
	fmt.Println("a",r.URL.Query().Get("image"))
	if InitStruct.Person.Genders == "h" {
		InitStruct.Person.Pants, err = strconv.Atoi(string(r.URL.Query().Get("image")[1]))
		if err != nil {
			fmt.Println("Erreur atoi", err.Error())
		}
		InitStruct.Person.Shirt, err = strconv.Atoi(string(r.URL.Query().Get("image")[0]))

		if err != nil {
			fmt.Println("Erreur atoi", err.Error())
		}

	} else {
		InitStruct.Person.Shirt, err = strconv.Atoi(string(r.URL.Query().Get("image")))

	}

	InitStruct.Persons = append(InitStruct.Persons, InitStruct.Person)
	InitStruct.EditJSON(InitStruct.Persons)

	http.Redirect(w, r, "/index", http.StatusMovedPermanently)
}

func Gender(w http.ResponseWriter, r *http.Request) {
	InitTemps.Temp.ExecuteTemplate(w, "gender", nil)

}
