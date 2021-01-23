package main

import "log"

func init() {
	configureWebserver()
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
