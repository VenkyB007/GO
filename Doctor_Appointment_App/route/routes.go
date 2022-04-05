package route

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vikasbolla/DoctorApplication/Doctor_application/service"
)

func Start() {
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/login", service.Login).Methods("POST")
	myRouter.HandleFunc("/getall", service.DoctorDetails).Methods("GET")
	myRouter.HandleFunc("/postavailability", service.PostAvailability).Methods("POST")
	myRouter.HandleFunc("/getsuitableappointment", service.GetSuitableAppointment).Methods("POST")
	log.Fatal((http.ListenAndServe(":9009", myRouter)))
}
