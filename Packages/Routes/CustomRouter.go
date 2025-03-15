package Routes

import (
	"ExpenceTraker/Helper"
	"ExpenceTraker/Packages/Controller"

	"github.com/gorilla/mux"
)

func CustomRouter(Router *mux.Router) {
	Router.HandleFunc(Helper.SignUpRoute, Controller.SignUp).Methods("GET")
	Router.HandleFunc(Helper.LoginInRoute, Controller.Login).Methods("GET")

	Router.HandleFunc(Helper.AddExpenseRoute, Controller.AddExpenceControl).Methods("POST")
	Router.HandleFunc(Helper.GetExpenseRoute, Controller.GetExpenseControl).Methods("GET")
	Router.HandleFunc(Helper.UpdateExpenseRoute, Controller.UpdateExpenseControl).Methods("PUT")
	Router.HandleFunc(Helper.RemoveExpenseRoute, Controller.DeleteExpenseControl).Methods("DELETE")
}
