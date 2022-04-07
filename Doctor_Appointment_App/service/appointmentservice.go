package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/VenkyB007/go/Doctor_Appointment_App/domain"
)

var respdoctor []domain.RespDoctor

var appointment domain.Appointment

func GetSuitableAppointment(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&appointment)
	Gettintvalues()
	var apptfstr string
	var appttstr string
	for i := 11; i < 13; i++ {
		apptfstr = apptfstr + string(appointment.From[i])
		appttstr = appttstr + string(appointment.To[i])
	}
	apptfrom, _ := strconv.Atoi(apptfstr)
	apptto, _ := strconv.Atoi(appttstr)
	for i := 0; i < len(Doctors); i++ {
		c := 0
		var respd domain.RespDoctor
		if apptfrom >= Slot[i].TimeInt[0].From && apptfrom < Slot[i].TimeInt[0].To {
			respd.DoctorId = Slot[i].Id
			respd.Fullname = Slot[i].Name
			c++
		} else if apptto < Slot[i].TimeInt[1].To && apptto > Slot[i].TimeInt[1].From {
			respd.DoctorId = Slot[i].Id
			respd.Fullname = Slot[i].Name
			c++
		}
		if c == 1 {
			respdoctor = append(respdoctor, respd)
		}
	}

	json.NewEncoder(w).Encode(respdoctor)
}
