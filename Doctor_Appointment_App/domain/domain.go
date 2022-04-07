package domain

import "github.com/dgrijalva/jwt-go"

//Doctor Details && Time Slots
type Doctor struct {
	DoctorId     string  `json:"doctorId"`
	Fullname     string  `json:"fullname"`
	Availability []Avail `json:"availability"`
}

type Avail struct {
	From string `json:"from"`
	To   string `json:"to"`
}

//Suitable Appointment Time Slots with Id
type Appointment struct {
	AppointmentId string `json:"appoinmentid"`
	From          string `json:"from"`
	To            string `json:"to"`
}

//Duplicate of Doctor Details && int Time Slots for decision making
type Doctordup struct {
	Id      string `json:"doctorid"`
	Name    string `json:"fullname"`
	TimeInt [2]TimeInt
}

type TimeInt struct {
	From int `json:"from"`
	To   int `json:"to"`
}

//Final Response with required details
type RespDoctor struct {
	DoctorId string `json:"doctorId"`
	Fullname string `json:"fullname"`
}

//Authentication

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
