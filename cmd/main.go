package main

import (
	"context"
	"log"
	"net/http"

	"github.com/bagusandrian/sawitpro/config"
	httpSawitProHandler "github.com/bagusandrian/sawitpro/handler/http"
	httpSawitProImpl "github.com/bagusandrian/sawitpro/handler/http/impl"
	"github.com/bagusandrian/sawitpro/resouce"
	"github.com/gorilla/mux"
)

func main() {
	ctx := context.Background()
	conf, err := config.New(ctx)
	if err != nil {
		log.Panicf("failed to init the config: %v", err)
	}
	router := mux.NewRouter()
	resource, err := resouce.InitResource(conf)
	if err != nil {
		log.Panicf("failed to init the resource: %v", err)
	}
	handler := httpSawitProImpl.New(conf, resource)
	RegisterRouter(router, handler)
	log.Printf("Server is running on port %s\n", conf.Server.HTTP.Address)
	log.Fatal(http.ListenAndServe(conf.Server.HTTP.Address, router))
}

func RegisterRouter(router *mux.Router, h httpSawitProHandler.Handler) {
	router.HandleFunc("/registration", h.Registration).Methods("POST")
	router.HandleFunc("/login", h.Login).Methods("POST")
	router.HandleFunc("/get_my_profile", h.GetProfile).Methods("GET")
	router.HandleFunc("/update_profile", h.UpdateProfile).Methods("POST")
}
