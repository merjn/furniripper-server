package main

import (
	"github.com/rs/zerolog/log"
)

func init() {
	log.Info().Msg("Habbo.ovh furniripper server")
	log.Info().Msg("Developed by https://github.com/merjn")

	configureWebserver()
}

func main() {
	if err := Run(); err != nil {
		log.Fatal().Msg(err.Error())
	}
}
