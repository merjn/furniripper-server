package main

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

// Run starts the application.
func Run() error {
	defer db.Close()
	log.Info().Msgf("Habbo furniripper server started on port %d", c.WebserverPort)

	listenTo := fmt.Sprintf(":%d", c.WebserverPort)
	if err := http.ListenAndServe(listenTo, mux); err != nil {
		return err
	}

	return nil
}
