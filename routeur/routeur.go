package routeur

import (
	"fmt"
	ctrl "koh-lanta/controller"
	"net/http"
	"os"
)

func InitServe() {
	/*Initialisation des routes
	http.HandleFunc("Route actuel, fonction activé")
	Lorsque on se situe sur une route la fonction associé va s'activé */
	http.HandleFunc("/index", ctrl.Index)
	http.HandleFunc("/add", ctrl.Add)
	http.HandleFunc("/add/treatment", ctrl.InitAdd)
	http.HandleFunc("/display", ctrl.Display)
	http.HandleFunc("/update", ctrl.Update)
	http.HandleFunc("/update/treatment", ctrl.InitUpdate)
	http.HandleFunc("/suppr", ctrl.Suppr)
	http.HandleFunc("/search", ctrl.Search)
	http.HandleFunc("/gender", ctrl.Gender)



	//renvoie sur la page d'erreur si la route n'est pas trouvée
	http.HandleFunc("/", ctrl.HandleError)

	//Pour relier les assets(img/fonts/css) aux templates
	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	//Le lien d'ou est lancé le serveur
	fmt.Println("(http://localhost:8080/index) - Server started on port:8080")
	fmt.Println("Si le navigateur ne c'est pas ouvrir tous seul, va y tous seul et tape  http://localhost:8080/index  dans ton navigateur préféré.")
	fmt.Println("Si tu veux arrêter le server fait un CTRL + C dans le terminal ")
	http.ListenAndServe("localhost:8080", nil)
	fmt.Println("Server closed")
}
