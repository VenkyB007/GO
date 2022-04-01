package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/", homepage)
	myRouter.HandleFunc("/doctors", doctorDetails).Methods("GET")
	myRouter.HandleFunc("/createdoc", createDoc).Methods("POST")
	log.Fatal((http.ListenAndServe(":10000", myRouter)))
}
