package main

import (
	"context"
	"log"
	"net/http"

	"github.com/bagusandrian/sawitpro/config"
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
	httpSawitProImpl.New(router, conf, resource)
	log.Printf("Server is running on port %s\n", conf.Server.HTTP.Address)
	log.Fatal(http.ListenAndServe(conf.Server.HTTP.Address, router))
}
