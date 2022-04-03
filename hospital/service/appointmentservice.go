package service

import (
	"encoding/json"
	"net/http"

	"github.com/VenkyB007/go/hospital/domain"
)

var Appointment domain.Appointment

func GetSuitableAppointment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&Appointment)

	for i := 0; i < len(domain.Doctors); i++ {
		for j := 0; j < 2; j++ {
			if domain.Doctors[i].Availability[j].From == Appointment.From && domain.Doctors[i].Availability[j].To == Appointment.To {
				var ResponseDoctor domain.RespDoctor
				ResponseDoctor.DoctorId = domain.Doctors[i].DoctorId
				ResponseDoctor.Fullname = domain.Doctors[i].Fullname

				json.NewEncoder(w).Encode(ResponseDoctor)
			}
		}
	}
}
