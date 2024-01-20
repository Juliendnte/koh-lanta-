package back

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// Fonction pour modifié le JSON
func EditJSON(ModifiedArticle []Personnage) {

	modifiedJSON, errMarshal := json.Marshal(ModifiedArticle)
	if errMarshal != nil {
		fmt.Println("Error encodage ", errMarshal.Error())
		return
	}

	// Écrire le JSON modifié dans le fichier
	err := os.WriteFile("JSON/bdd.json", modifiedJSON, 0644)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture du fichier JSON modifié:", err)
		return
	}

}

// Fonction pour mettre le JSON dans une struct
func ReadJSON() ([]Personnage, error) {
	jsonFile, err := os.ReadFile("JSON/bdd.json")
	if err != nil {
		fmt.Println("Error reading", err.Error())
	}

	var jsonData []Personnage
	err = json.Unmarshal(jsonFile, &jsonData)
	return jsonData, err
}

// Fonction pour savoir si l'id existe déjà
func IdAlreadyExists(nb int) bool {
	for i := 0; i < len(Persons); i++ {
		if Persons[i].Id == nb {
			return true
		}
	}
	return false
}

// Fonction pour générer un Id disponible
func GenerateID() int {
	var Id int = rand.Intn(100)
	if IdAlreadyExists(Id) {
		return GenerateID()
	}
	return Id
}

// Fonction pour rechercher
func Search(word string, s string) bool {
	return strings.Contains(strings.ToLower(word), strings.ToLower(s))
}
