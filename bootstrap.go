package main

import (
	"github.com/merjn/furniripper-server/handlers"
	"github.com/merjn/furniripper-server/middleware"
	"log"
	"net/http"
)

var mux *http.ServeMux

func configureWebserver() {
	mux = http.NewServeMux()

	configureAddFurniHandler()

	log.Println("Webserver mux configured")
}

func configureAddFurniHandler() {
	addFurniHandler := handlers.AddFurniHandler
	jwtTokenMiddleware := middleware.AuthorizeJwtToken(addFurniHandler)
	mux.HandleFunc("/add_furni", jwtTokenMiddleware)

	log.Println("Add furni handler configured")
}
