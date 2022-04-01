package domain

type Avail struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type Doctor struct {
	DoctorId     int     `json:"doctorId"`
	Fullname     string  `json:"fullname"`
	Availability []Avail `json:"availability"`
}

var Doctors []Doctor
