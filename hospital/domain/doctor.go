package domain

type Avail struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type Doctor struct {
	DoctorId     string  `json:"doctorId"`
	Fullname     string  `json:"fullname"`
	Availability []Avail `json:"availability"`
}

type RespDoctor struct {
	DoctorId string `json:"doctorId"`
	Fullname string `json:"fullname"`
}

var Doctors []Doctor
