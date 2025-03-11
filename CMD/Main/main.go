package main

import (
	"ExpenceTraker/Helper"
	"ExpenceTraker/Packages/Routes"
	Utility "ExpenceTraker/Packages/Utilities"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	Utility.InitialiseDatabaseConnection()

	Utility.InitialiseRedisConn()

	MuxRouter := mux.NewRouter()

	Routes.CustomRouter(MuxRouter)

	http.Handle(Helper.BaseRoute, MuxRouter)

	log.Fatal(http.ListenAndServe("localhost:9030", MuxRouter))

	defer Utility.TerminateDatabaseConnection()
}
