package main

import (
	"database/sql"
	"net/http"
	"os"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"github.com/merjn/furniripper-server/config"
	"github.com/merjn/furniripper-server/furni"
	"github.com/merjn/furniripper-server/handlers"
	"github.com/merjn/furniripper-server/service"
	"github.com/rs/zerolog/log"
	"github.com/tkanos/gonfig"
)

var mux *http.ServeMux
var c config.Config
var authMiddleware *jwtmiddleware.JWTMiddleware
var db *sql.DB

func createDatabase() {
	d, err := sql.Open("mysql", c.ConnectionString)
	if err != nil {
		log.Fatal().Msgf("unable to open mysql connection: %s", err.Error())
	}

	if err := d.Ping(); err != nil {
		log.Fatal().Msgf("unable to ping database: %s", err.Error())
	}

	db = d
}

func configureWebserver() {
	mux = http.NewServeMux()

	configureAddFurniHandler()

	log.Info().Msg("Webserver mux configured")
}

func configureJwtToken() {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		log.Fatal().Msg("JWT_SECRET not found in environment")
	}
	authMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (i interface{}, e error) {
			return []byte(secretKey), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	/*r := jwt.NewWithClaims(authMiddleware.Options.SigningMethod, jwt.MapClaims{
		"domain": "*",
	})

	tokenStr, err := r.SignedString([]byte(secretKey))
	if err != nil {
		panic(err)
	}

	fmt.Println(tokenStr)*/
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
	Arcturus := &furni.ArcturusAdder{
		DB: db,
	}

	FurniService := &service.Furni{
		Config: c,
		Adder:  Arcturus,
	}

	addFurniHandler := handlers.AddFurniHandler{
		FurniService: FurniService,
	}

	mux.Handle("/add_furni", authMiddleware.Handler(http.HandlerFunc(addFurniHandler.Handle)))

	log.Info().Msg("Add furni handler configured")
}
