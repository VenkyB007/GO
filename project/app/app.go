package main

import (
	"log"
	"net/http"
)

func Start() {
	http.HandleFunc("/greet", homepage)
	http.HandleFunc("/get_availability", get_availability)

	log.Fatal(http.ListenAndServe("localhost:8900", nil))
}
