package service

import (
	"net/http"
)

type student interface {
	GetStudentDetails(writer http.ResponseWriter, request *http.Request)
}

func GetStudentDetails(writer http.ResponseWriter, request *http.Request) {

}
