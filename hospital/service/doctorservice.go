package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func doctorDetails(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Doctors)
	//fmt.Fprintf(w, "doctordetails endpoint hit")
	fmt.Println(Doctors)
}

func createDoc(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var doctor Doctor
	json.NewDecoder(r.Body).Decode(&doctor)
	Doctors = append(Doctors, doctor)
	json.NewEncoder(w).Encode(doctor)
}
