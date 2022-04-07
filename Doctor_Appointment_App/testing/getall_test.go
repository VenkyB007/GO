package testing

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/VenkyB007/go/Doctor_Appointment_App/service"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

//////////////////Testing Get All Doctor Details//////////////////
func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/getall", service.DoctorDetails).Methods("GET")
	return router
}

func TestDoctorDetails(t *testing.T) {
	request, _ := http.NewRequest("GET", "/getall", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "*OK* response is expected")
}
