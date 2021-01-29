package main

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

// Run starts the application.
func Run() error {
	log.Info().Msgf("Habbo furniripper server started on port %d", config.WebserverPort)

	listenTo := fmt.Sprintf(":%d", config.WebserverPort)
	if err := http.ListenAndServe(listenTo, mux); err != nil {
		return err
	}

	return nil
}
