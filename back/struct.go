package back

type Personnage struct {
	Id      int    `json:"id"`
	Genders string `json:"genders"`
	Clothes int `json:"clothes"`
	Name    string `json:"name"`
}

var Person Personnage
var Persons []Personnage
