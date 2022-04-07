package route

import (
	"log"
	"net/http"

	"github.com/VenkyB007/go/Doctor_Appointment_App/service"
	"github.com/gorilla/mux"
)

func Start() {
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/login", service.Login).Methods("POST")
	myRouter.HandleFunc("/getall", service.DoctorDetails).Methods("GET")
	myRouter.HandleFunc("/postavailability", service.PostAvailability).Methods("POST")
	myRouter.HandleFunc("/getsuitableappointment", service.GetSuitableAppointment).Methods("POST")
	log.Fatal((http.ListenAndServe(":9009", myRouter)))
}
