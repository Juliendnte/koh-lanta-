package back

type Personnage struct {
	Id      int    `json:"id"`
	Genders string   `json:"genders"`
	Shirt   int    `json:"shirt"`
	Pants   int    `json:"pants"`
	Name    string `json:"name"`
}

var Person Personnage
var Persons []Personnage
var LstIDSuppr []int
