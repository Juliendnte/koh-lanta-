package main

import (
	r "koh-lanta/routeur" //Route vers mes routes
	t "koh-lanta/temps"   //Route vers mes templates
)

func main() {
	t.InitTemplate() //Initialise mes templates
	r.InitServe()    //Initialise mes routes / assets et lance le serveur
}