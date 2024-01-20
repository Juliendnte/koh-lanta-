package controller

import (
	"fmt"
	InitStruct "koh-lanta/back"
	InitTemps "koh-lanta/temps"
	"net/http"
	"os"
)

//Execute le template qui affiche tous les personnages
func Display(w http.ResponseWriter, r *http.Request) {
	InitStruct.Persons,err = InitStruct.ReadJSON()
	if err != nil {
		fmt.Println("Erreur read", err.Error())
	}
	
	InitTemps.Temp.ExecuteTemplate(w, "display", InitStruct.Persons)
}

//Execute le template error 404
func HandleError(w http.ResponseWriter, r *http.Request) {
	InitTemps.Temp.ExecuteTemplate(w, "error", nil)
}

// Fonction de la page des résultats d'une recherche
func Search(w http.ResponseWriter, r *http.Request) {
	InitStruct.Persons, err = InitStruct.ReadJSON() //Met le fichier JSON dans ma struct
	if err != nil {
		fmt.Println("Error encodage ", err.Error())
		os.Exit(1)
	}

	queryName := r.URL.Query().Get("text") //Récupére le text donné dans le Query string
	var lstSearch []InitStruct.Personnage

	for _, c := range InitStruct.Persons {
		if InitStruct.Search(c.Name, queryName) {
			lstSearch = append(lstSearch, c) //Met tous les personnages résultant de la recherche
		}
	}

	if len(lstSearch) == 0 { //Oui je triche
		none := InitStruct.Personnage{Name: "Nothing", Id: 0, Genders: "", Clothes: 0}
		lstSearch = append(lstSearch, none)
	}

	//execute le template Search
	InitTemps.Temp.ExecuteTemplate(w, "Search", lstSearch)
}
