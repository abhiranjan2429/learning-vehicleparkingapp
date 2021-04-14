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
	w.Write([]byte(`Hi this is the home page for vehicle entry go to /entry and the json body needed is \n
	{
		"vehiclenumber":"XXXXXX"
	}`))
}
