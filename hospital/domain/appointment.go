package domain

type Appointment struct {
	AppointmentId string `json:"appointmentid"`
	From          string `json:"from"`
	To            string `json:"to"`
}
