package main

import (
	"net/http"

	"github.com/merjn/furniripper-server/config"
	"github.com/merjn/furniripper-server/handlers"
	"github.com/merjn/furniripper-server/middleware"
	"github.com/merjn/furniripper-server/service"
	"github.com/rs/zerolog/log"
	"github.com/tkanos/gonfig"
)

var mux *http.ServeMux
var c config.Config

func configureWebserver() {
	mux = http.NewServeMux()

	configureAddFurniHandler()

	log.Info().Msg("Webserver mux configured")
}

func setConfig() {
	var conf config.Config
	err := gonfig.GetConf("config.json", &conf)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	c = conf
}

func configureAddFurniHandler() {
	FurniService := &service.Furni{
		Config: c,
	}

	addFurniHandler := handlers.AddFurniHandler{
		FurniService: FurniService,
	}

	jwtTokenMiddleware := middleware.AuthorizeJwtToken(addFurniHandler.Handle)
	mux.HandleFunc("/add_furni", jwtTokenMiddleware)

	log.Info().Msg("Add furni handler configured")
}
