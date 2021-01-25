package main

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

// Run starts the application.
func Run() error {
	log.Info().Msg("Habbo furniripper server started on port 3000")

	if err := http.ListenAndServe(":3000", mux); err != nil {
		return err
	}

	return nil
}
