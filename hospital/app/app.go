package app

import (
	"log"
	"net/http"

	"github.com/VenkyB007/go/hospital/service"
	"github.com/gorilla/mux"
)

func Start() {

	myRouter := mux.NewRouter()

	myRouter.HandleFunc("/getall", service.DoctorDetails).Methods("GET")
	myRouter.HandleFunc("/postavailability", service.PostAvailability).Methods("POST")
	myRouter.HandleFunc("/getsuitableappointment", service.GetSuitableAppointment).Methods("POST")
	log.Fatal((http.ListenAndServe(":9019", myRouter)))
}
