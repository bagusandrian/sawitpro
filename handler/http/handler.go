package http

import "net/http"

//go:generate mockery --name=Handler --filename=mock_handler.go --inpackage
type Handler interface {
	Registration(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	GetProfile(w http.ResponseWriter, r *http.Request)
	UpdateProfile(w http.ResponseWriter, r *http.Request)
}
