package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/VenkyB007/go/hospital/domain"
)

var doctor domain.Doctor
var Doctors []domain.Doctor

// func DoctorDetails(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Add("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(domain.Doctors)

// 	fmt.Println(domain.Doctors)
// }

func DoctorDetails(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Doctors)

	fmt.Println(Doctors)
}

func PostAvailability(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	json.NewDecoder(r.Body).Decode(&doctor)
	json.NewEncoder(w).Encode(doctor)
	Doctors = append(Doctors, doctor)

	//fmt.Fprintf(w, "Successfully Posted")

}
