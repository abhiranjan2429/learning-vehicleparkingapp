package controller

import (
	"net/http"
)

type (
	home struct {
	}
)

func (e home) registerRoutes() {
	http.HandleFunc("/home", e.handlehome)
	http.HandleFunc("/", e.handlehome)
}
func (e home) handlehome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`Hi this is the home page for billing service go to
	 /billing for billcalculation
	 /prktime for time calculation
	  and the json body needed is for both billing and prktime is 
	{
		"parkingid":"XXXXXX"
	}`))
}
