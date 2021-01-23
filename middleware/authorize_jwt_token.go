package middleware

import (
	"log"
	"net/http"
)

func AuthorizeJwtToken(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Println("Authorizing JWT token...")
		next.ServeHTTP(w, req)
	})
}