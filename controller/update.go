package controller

import (
	"fmt"
	InitStruct "koh-lanta/back"
	InitTemps "koh-lanta/temps"
	"net/http"
	"os"
	"strconv"
)

var Id int
var Genderr string

func Update(w http.ResponseWriter, r *http.Request) {
	queryID, errId := strconv.Atoi(r.URL.Query().Get("id"))
	if errId != nil {
		fmt.Println("Error ID ", errId.Error())
		os.Exit(1)
	}

	//Prend le blog à l'id donné
	for _, i := range InitStruct.Persons {
		if i.Id == queryID {
			InitStruct.Person = i
			Genderr = i.Genders
			break
		}
	}
	fmt.Println(InitStruct.Person, "up")
	Id = queryID

	if InitStruct.Person.Genders == "h" {
		InitTemps.Temp.ExecuteTemplate(w, "update", InitStruct.Person)
	} else if InitStruct.Person.Genders == "f" {
		InitTemps.Temp.ExecuteTemplate(w, "updatef", InitStruct.Person)

	}
}

func InitUpdate(w http.ResponseWriter, r *http.Request) {
	InitStruct.Persons, err = InitStruct.ReadJSON()
	if err != nil {
		fmt.Println("Erreur read", err.Error())
	}

	//Prend le blog à l'id donné
	var c int
	fmt.Println("IN IN")
	for _, i := range InitStruct.Persons {
		fmt.Println("Update", i.Id, Id)
		if i.Id == Id {
			InitStruct.Person.Id = i.Id
			InitStruct.Person.Genders = Genderr
			InitStruct.Person.Name = r.FormValue("name")
			InitStruct.Person.Clothes, err = strconv.Atoi(r.URL.Query().Get("image"))
			if err != nil {
				fmt.Println("Erreur atoi", err.Error())
			}
			break
		}
		c++
	}
	fmt.Println( InitStruct.Person)
	InitStruct.Persons[c] = InitStruct.Person
	InitStruct.EditJSON(InitStruct.Persons)
	http.Redirect(w, r, "/display", http.StatusMovedPermanently)
}

func Suppr(w http.ResponseWriter, r *http.Request) {
	InitStruct.Persons, err = InitStruct.ReadJSON() //Met le fichier JSON dans ma struct
	if err != nil {
		fmt.Println("Error encodage ", err.Error())
		os.Exit(1)
	}

	queryID, errId := strconv.Atoi(r.URL.Query().Get("id")) //Récupére l'Id donné dans le Query string et le met en int
	if errId != nil {
		fmt.Println("Error ID ", errId.Error())
		os.Exit(1)
	}

	for i, c := range InitStruct.Persons {
		if c.Id == queryID {
			InitStruct.Persons = append(InitStruct.Persons[:i], InitStruct.Persons[i+1:]...) //Supprime de la liste des blogs
			break
		}
	}
	InitStruct.EditJSON(InitStruct.Persons) //Met les données dans le JSON
	http.Redirect(w, r, "/display", http.StatusMovedPermanently)
}
