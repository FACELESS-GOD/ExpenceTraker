package Routes

import (
	"ExpenceTraker/Helper"
	"ExpenceTraker/Packages/Controller"

	"github.com/gorilla/mux"
)

func CustomRouter(Router *mux.Router) {
	Router.HandleFunc(Helper.SignUpRoute, Controller.SignUp).Methods("GET")
	Router.HandleFunc(Helper.LoginInRoute, Controller.Login).Methods("GET")
}
