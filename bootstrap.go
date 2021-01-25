package main

import (
	"net/http"

	"github.com/merjn/furniripper-server/handlers"
	"github.com/merjn/furniripper-server/middleware"
	"github.com/merjn/furniripper-server/service"
	"github.com/rs/zerolog/log"
)

var mux *http.ServeMux

func configureWebserver() {
	mux = http.NewServeMux()

	configureAddFurniHandler()

	log.Info().Msg("Webserver mux configured")
}

func configureAddFurniHandler() {
	addFurniHandler := handlers.AddFurniHandler{
		FurniService: new(service.Furni),
	}

	jwtTokenMiddleware := middleware.AuthorizeJwtToken(addFurniHandler.Handle)
	mux.HandleFunc("/add_furni", jwtTokenMiddleware)

	log.Info().Msg("Add furni handler configured")
}
