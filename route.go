package routes

import (
	"log"
	"github.com/gorilla/mux"
	"manger/controllers"
	"net/http"
)

   func  Register(){

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/smiles/{name}", controllers.Api).Methods("POST")

	log.Fatal(http.ListenAndServe(":8081", myRouter))

	}