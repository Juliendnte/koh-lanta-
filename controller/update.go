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
			break
		}
	}

	Id = queryID

	InitTemps.Temp.ExecuteTemplate(w, "update", InitStruct.Person)
}

func InitUpdate(w http.ResponseWriter, r *http.Request) {
	InitStruct.Persons, err = InitStruct.ReadJSON()
	if err != nil {
		fmt.Println("Erreur atoi", err.Error())
	}

	//Prend le blog à l'id donné
	for _, i := range InitStruct.Persons {
		if i.Id == Id {
			InitStruct.Persons[Id-1].Genders = r.FormValue("sexe")
			InitStruct.Persons[Id-1].Name = r.FormValue("name")
			InitStruct.Persons[Id-1].Pants, err = strconv.Atoi(r.FormValue("pants"))
			if err != nil {
				fmt.Println("Erreur atoi", err.Error())
			}
			InitStruct.Persons[Id-1].Shirt, err = strconv.Atoi(r.FormValue("shirt"))
			if err != nil {
				fmt.Println("Erreur atoi", err.Error())
			}
			break
		}
	}

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

	queryID-- //Pour que l'index commence à 0
	for _, c := range InitStruct.Persons {
		if c.Id == queryID {
			InitStruct.Persons = append(InitStruct.Persons[:queryID], InitStruct.Persons[queryID+1:]...) //Supprime de la liste des blogs
			queryID++
			InitStruct.LstIDSuppr = append(InitStruct.LstIDSuppr, queryID)
			break
		}
	}
	InitStruct.EditJSON(InitStruct.Persons) //Met les données dans le JSON
	http.Redirect(w, r, "/display", http.StatusMovedPermanently)
}
