package main

import (
	"net/http"
)

// Run starts the application.
func Run() error {
	if err := http.ListenAndServe(":3000", mux); err != nil {
		return err
	}

	return nil
}